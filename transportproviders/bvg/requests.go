package bvg

import (
	"time"
)

type baseRequest struct {
	Lang   string `json:"lang"`
	Ext    string `json:"ext"`
	Ver    string `json:"ver"`
	Auth   auth   `json:"auth"`
	Client client `json:"client"`
}

type auth struct {
	Type string `json:"type"`
	Aid  string `json:"aid"`
}

type client struct {
	Type string `json:"type"`
	ID   string `json:"id"`
	V    string `json:"v"`
	Name string `json:"name"`
}

type StationRequestPayload struct {
	baseRequest
	SvcReqL []LocDetailsSvcReq `json:"svcReqL"`
}

type DepartureRequestPayload struct {
	baseRequest
	SvcReqL []StationBoardSvcReq `json:"svcReqL"`
}

type StationBoardSvcReq struct {
	Meth string          `json:"meth"`
	Req  StationBoardReq `json:"req"`
}

type LocDetailsSvcReq struct {
	Meth string        `json:"meth"`
	Req  LocDetailsReq `json:"req"`
}

type LocDetailsReq struct {
	LocL []LocationReq `json:"locL"`
}

type LocationReq struct {
	Type string `json:"type"`
	Lid  string `json:"lid"`
}

type StationBoardReq struct {
	Type     string      `json:"type"`
	Date     string      `json:"date"`
	Time     string      `json:"time"`
	StbLoc   LocationReq `json:"stbLoc"`
	JnyFltrL []JnyFltr   `json:"jnyFltrL"`
	Dur      int         `json:"dur"` // Duration for which to return departures
}

type JnyFltr struct {
	Type  string `json:"type"`
	Mode  string `json:"mode"`
	Value uint8  `json:"value"` // bitmask sum to filter products
}

type DepartureRequestPayloadParams struct {
	when           time.Time
	stationID      string
	productsFilter uint8
	duration       int
}

func CreateStationRequestPayload(id string) StationRequestPayload {
	return StationRequestPayload{
		baseRequest{
			Lang: "en",
			Ext:  "BVG.1",
			Ver:  "1.44",
			Auth: auth{
				Type: "AID",
				Aid:  "YoJ05NartnanEGCj",
			},
			Client: client{
				Type: "IPA",
				ID:   "BVG",
				V:    "6020000",
				Name: "FahrInfo",
			},
		},
		[]LocDetailsSvcReq{
			{
				Meth: "LocDetails",
				Req: LocDetailsReq{
					LocL: []LocationReq{
						{Type: "S", Lid: "A=1@L=" + id + "@"},
					},
				},
			},
		},
	}
}

func CreateDepartureRequestPayload(params DepartureRequestPayloadParams) DepartureRequestPayload {

	return DepartureRequestPayload{
		baseRequest{
			Lang: "en",
			Ext:  "BVG.1",
			Ver:  "1.44",
			Auth: auth{
				Type: "AID",
				Aid:  "YoJ05NartnanEGCj",
			},
			Client: client{
				Type: "IPA",
				ID:   "BVG",
				V:    "6020000",
				Name: "FahrInfo",
			},
		},
		[]StationBoardSvcReq{
			{
				Meth: "StationBoard",
				Req: StationBoardReq{
					Type: "DEP",
					Date: params.when.Format(dateLayout),
					Time: params.when.Format(timeLayout),
					StbLoc: LocationReq{
						Type: "S",
						Lid:  "A=1@L=" + params.stationID + "@",
					},
					JnyFltrL: []JnyFltr{
						{
							Type:  "PROD",
							Mode:  "INC",
							Value: params.productsFilter,
						},
					},
					Dur: params.duration,
				},
			},
		},
	}
}
