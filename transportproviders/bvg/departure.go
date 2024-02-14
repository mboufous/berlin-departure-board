package bvg

import (
	"errors"
	"fmt"
	"github.com/mboufous/berlin-departure-board/hafas"
	"io"
	"log/slog"
	"math"
	"net/http"
	"time"
)

func (p *APIProvider) NewDepartureRequest(params *hafas.DepartureParams) (*http.Request, error) {
	payload := CreateDepartureRequestPayload(DepartureRequestPayloadParams{
		stationID: params.Station,
		duration:  params.MaxDeparturesDurationMinutes,
	})
	return p.newRequest(payload)
}

// TODO: process canceled departures
func (p *APIProvider) ParseDepartureResponse(body io.ReadCloser) (*hafas.DepartureBoard, error) {
	defer body.Close()

	apiResult, err := p.ParseBaseResponse(body)
	if err != nil {
		return nil, err
	}

	if apiResult.Meth != "StationBoard" {
		return nil, errors.New("wrong api call for departures")
	}

	return &hafas.DepartureBoard{
		Lines:   p.convertLines(apiResult.Res),
		Remarks: p.convertStationRemarks(apiResult.Res.Common.HimL),
	}, nil
}

func (p *APIProvider) convertDirectionRemarks(assignedRemarks []MsgList, remarks []RemMessage) []hafas.Remark {
	if len(assignedRemarks) == 0 || len(remarks) == 0 {
		return nil
	}
	var directionRemarks []hafas.Remark
	for i, remark := range assignedRemarks {
		if remark.Type == "REM" {
			directionRemarks = append(directionRemarks, hafas.Remark{
				Header: fmt.Sprintf("Remark%d", i),
				Body:   remarks[remark.RemX].TxtN,
			})
		}
	}
	return directionRemarks
}

func (p *APIProvider) convertLines(source *SvcResData) []hafas.Line {
	var lines []hafas.Line
	processedLines := make(map[string]int)
	lastLineIndex := -1

	for _, journey := range source.JnyL {
		product := source.Common.ProdL[journey.StbStop.DProdX].ProdCtx
		when, _, err := p.populateDepartureTime(&journey)
		if err != nil {
			slog.Error("error populating departure time", slog.String("Error", err.Error()))
			continue
		}

		lineID := product.Line
		lineIndex, exist := processedLines[lineID]
		if !exist {
			lastLineIndex++
			processedLines[lineID] = lastLineIndex
			lines = append(lines, hafas.Line{
				Product: hafas.Product{
					Name: product.Line,
					Type: product.CatOutS,
				},
				Directions: []hafas.Direction{},
			})
			lineIndex = lastLineIndex
		}

		directionIndex := -1
		for i, direction := range lines[lineIndex].Directions {
			if journey.DirTxt == direction.Name {
				directionIndex = i
				break
			}
		}
		if directionIndex == -1 {
			directionIndex = len(lines[lineIndex].Directions)
			lines[lineIndex].Directions = append(lines[lineIndex].Directions, hafas.Direction{
				Name:    journey.DirTxt,
				Remarks: p.convertDirectionRemarks(journey.MsgL, source.Common.RemL),
				Departures: []hafas.Departure{
					{When: when},
				},
			})
		} else {
			lines[lineIndex].Directions[directionIndex].Departures = append(lines[lineIndex].Directions[directionIndex].Departures, hafas.Departure{When: when})
		}
	}

	return lines
}

func (p *APIProvider) convertStationRemarks(source []HimMessage) []hafas.Remark {
	var remarks []hafas.Remark
	for _, remark := range source {
		remarks = append(remarks, hafas.Remark{
			Header: remark.Head,
			Body:   remark.Text,
		})
	}
	return remarks
}

// TODO: ignore invalid date times (negatives)
func (p *APIProvider) populateDepartureTime(journey *Journey) (time.Time, int, error) {
	plannedDepartureTimeRaw := journey.StbStop.DTimeS
	newDepartureTimeRaw := journey.StbStop.DTimeR
	departureDate := journey.Date

	// Parse the planned departure time
	departureTime, err := p.parseDepartureTime(plannedDepartureTimeRaw, departureDate)
	if err != nil {
		return time.Time{}, 0, fmt.Errorf("planned departure time parsing failed: %w", err)
	}

	delay := 0
	if p.departureDelayed(journey) {
		newDepartureTime, err := p.parseDepartureTime(newDepartureTimeRaw, departureDate)
		if err != nil {
			return time.Time{}, 0, fmt.Errorf("new departure time parsing failed: %w", err)
		}

		delay = int(math.Round(newDepartureTime.Sub(departureTime).Minutes()))
		departureTime = newDepartureTime
	}

	return departureTime, delay, nil
}

func (p *APIProvider) parseDepartureTime(timeStr, dateStr string) (time.Time, error) {
	// Load the Germany time zone
	location, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to load Germany time zone: %w", err)
	}

	var offsetDays time.Duration
	adjustedTime := timeStr

	if p.isDayOffsetPresent(timeStr) {
		adjustedTime = timeStr[2:]
		offsetDays, err = time.ParseDuration(fmt.Sprintf("%sh", timeStr[0:2]))
		if err != nil {
			return time.Time{}, fmt.Errorf("failed to parse time offset: %w", err)
		}
	}
	baseDate, err := time.ParseInLocation(fullDateLayout, dateStr+adjustedTime, location)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse date: %w", err)
	}

	return baseDate.Add(offsetDays * 24), nil
}

func (p *APIProvider) isDayOffsetPresent(timeStr string) bool {
	return len(timeStr) > len(timeLayout)
}

func (p *APIProvider) departureDelayed(journey *Journey) bool {
	return journey.StbStop.DTimeR != "" && journey.StbStop.DTimeS != journey.StbStop.DTimeR
}
