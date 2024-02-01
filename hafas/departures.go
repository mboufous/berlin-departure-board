package hafas

import (
	"context"
	"time"
)

type DeparturesService Service

type Departure struct {
	Stop        Station
	Direction   string
	Delay       uint
	When        time.Time // 'when' field includes the current delay
	PlannedWhen time.Time
	// Line Line
	//Remarks Remarks
	//
}

func (s *DeparturesService) Get(ctx context.Context, params any) (*Departure, error) {
	if ctx == nil {
		return nil, errNilContext
	}

	req, err := s.client.NewDepartureRequest()
}
