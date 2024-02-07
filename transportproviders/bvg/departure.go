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

func (p *Provider) NewDepartureRequest(params *hafas.DepartureParams) (*http.Request, error) {
	payload := CreateDepartureRequestPayload(DepartureRequestPayloadParams{
		stationID:      params.Station.ID,
		duration:       params.DurationMinutes,
		productsFilter: params.ProductsFilter,
	})
	return p.newRequest(payload)
}

// TODO: process canceled departures
func (p *Provider) ParseDepartureResponse(body io.ReadCloser, showRemarks bool) (*hafas.DepartureBoard, error) {
	defer body.Close()

	d := &hafas.DepartureBoard{}
	apiResult, err := p.ParseBaseResponse(body)
	if err != nil {
		return nil, err
	}

	if apiResult.Meth != "StationBoard" {
		return nil, errors.New("wrong api call for departures")
	}

	if showRemarks {
		d.Remarks = p.convertRemarks(apiResult.Res.Common.HimL)
	}
	d.Departures = p.convertDeparture(apiResult.Res)

	return d, nil
}

func (p *Provider) convertRemarks(source []HimMessage) []hafas.Remark {
	var remarks []hafas.Remark
	for _, remark := range source {
		remarks = append(remarks, hafas.Remark{
			Header: remark.Head,
			Body:   remark.Text,
		})
	}
	return remarks
}

func (p *Provider) convertDeparture(source *SvcResData) []hafas.Departure {
	var departures []hafas.Departure
	for _, journey := range source.JnyL {
		stop := source.Common.LocL[journey.StbStop.LocX]
		line := source.Common.ProdL[journey.StbStop.DProdX].ProdCtx.Line
		when, delay, err := p.populateDepartureTime(journey)

		if err != nil {
			slog.Warn("departure time couldn't be converted.", slog.Any("Err", err))
			continue
		}

		departures = append(departures, hafas.Departure{
			Stop: hafas.Station{
				ID:   stop.ExtId,
				Name: stop.Name,
			},
			Direction: journey.DirTxt,
			When:      when,
			Delay:     delay,
			Line: hafas.Line{
				Name: line,
			},
		})
	}
	return departures
}

// TODO: ignore invalid date times (negatives)
func (p *Provider) populateDepartureTime(journey Journey) (time.Time, int, error) {
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

func (p *Provider) parseDepartureTime(timeStr, dateStr string) (time.Time, error) {
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

func (p *Provider) isDayOffsetPresent(timeStr string) bool {
	return len(timeStr) > len(timeLayout)
}

func (p *Provider) departureDelayed(journey Journey) bool {
	return journey.StbStop.DTimeR != "" && journey.StbStop.DTimeS != journey.StbStop.DTimeR
}
