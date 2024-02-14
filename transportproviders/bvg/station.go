package bvg

import (
	"errors"
	"github.com/mboufous/berlin-departure-board/hafas"
	"io"
	"net/http"
)

func (p *APIProvider) NewStationRequest(params *hafas.StationParams) (*http.Request, error) {
	payload := CreateStationRequestPayload(params.StationID)
	return p.newRequest(payload)
}

func (p *APIProvider) ParseStationResponse(body io.ReadCloser) (*hafas.Station, error) {
	defer body.Close()
	apiResult, err := p.ParseBaseResponse(body)
	if err != nil {
		return nil, err
	}

	if apiResult.Meth != "LocDetails" {
		return nil, errors.New("wrong api call for station details")
	}

	station := p.convertStation(apiResult.Res)

	return station, nil
}

func (p *APIProvider) convertStation(source *SvcResData) *hafas.Station {
	stationData := &hafas.Station{
		ID:       source.LocL[0].ExtId,
		Name:     source.LocL[0].Name,
		Products: p.populateProducts(source.Common.ProdL),
	}
	return stationData
}

func (p *APIProvider) populateProducts(products []ProductReq) []hafas.Product {
	var hafasProducts []hafas.Product

	for _, product := range products {
		hafasProducts = append(hafasProducts, hafas.Product{
			Type: product.ProdCtx.CatOutS,
			Name: product.NameS,
		})
	}

	return hafasProducts
}
