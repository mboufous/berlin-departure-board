package providers

import (
	"context"
	"net/http"

	"github.com/mboufous/berlin-departure-board/internal"
)

const VBBProviderName = "VBB"

type VBBProvider struct {
	internal.BaseProvider
}

func NewVBBProvider(client *http.Client) *VBBProvider {
	return &VBBProvider{
		internal.BaseProvider{
			Client: client,
		},
	}
}

func (p *VBBProvider) GetDepartures(ctx context.Context, stationID string) ([]internal.Departure, error) {
	return []internal.Departure{
		{Line: "VBB Line"},
	}, nil
}

func (p *VBBProvider) HealthCheck() (internal.APIProviderStatus, error) {
	return internal.APIProviderStatus{
		Up: true,
	}, nil
}

func (p *VBBProvider) Name() string {
	return VBBProviderName
}

func (p *VBBProvider) String() string {
	return p.Name()
}
