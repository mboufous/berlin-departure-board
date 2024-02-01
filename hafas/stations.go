package hafas

import (
	"context"
)

type StationService Service

//const (
//	BitMaskSuburban = 1 << iota
//	BitMaskSubway
//	BitMaskTram
//	BitMaskBus
//	BitMaskFerry
//	BitMaskExpress
//	BitMaskRegional
//)
//
//const (
//	ProductTypeSuburban ProductType = "suburban"
//	ProductTypeSubway   ProductType = "subway"
//	ProductTypeTram     ProductType = "tram"
//	ProductTypeBus      ProductType = "bus"
//	ProductTypeFerry    ProductType = "ferry"
//	ProductTypeExpress  ProductType = "express"
//	ProductTypeRegional ProductType = "regional"
//)

type Station struct {
	ID       string
	Name     string
	Products []Product
}

type Product struct {
	Name string
	Type string
}

func (s *StationService) Get(ctx context.Context, params any) (*Station, error) {
	if ctx == nil {
		return nil, errNilContext
	}

	req, err := s.client.NewDepartureRequest(params)
	if err != nil {
		return nil, err
	}

	station, err := s.client.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	return station, nil
}
