package bvg

import (
	"fmt"

	"github.com/mboufous/berlin-departure-board/internal"
)

type BaseRequest struct {
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
	L    string `json:"l"`
}

type DepartureBoardRequest struct {
	BaseRequest
	SvcReqL []departureBoardSvcRequest `json:"svcReqL"`
}

type departureBoardSvcRequest struct {
	Meth string            `json:"meth"`
	Req  departureBoardReq `json:"req"`
}

type departureBoardReq struct {
	Type     string      `json:"type"`
	StbLoc   locationReq `json:"stbLoc"`
	JnyFltrL []jnyFltr   `json:"jnyFltrL"`
	// Dur      int         `json:"dur"` // Duration for which to return departures
	Sort string `json:"sort"`
	// MaxJny   int         `json:"maxJny"`
	// Date     string      `json:"date"`
	// Time     string      `json:"time"`
}

type jnyFltr struct {
	Type  string `json:"type"`
	Mode  string `json:"mode"`
	Value uint8  `json:"value"` // bitmask sum to filter products
}

type ServerStatusRequest struct {
	BaseRequest
	SvcReqL []serverStatusSvcRequest `json:"svcReqL"`
}

type serverStatusSvcRequest struct {
	Meth string `json:"meth"`
}

type StationRequest struct {
	BaseRequest
	SvcReqL []locDetailsSvcRequest `json:"svcReqL"`
}

type locDetailsSvcRequest struct {
	Meth string        `json:"meth"`
	Req  locDetailsReq `json:"req"`
}

type locDetailsReq struct {
	GetHIM        bool          `json:"getHIM"`
	GetProducts   bool          `json:"getProducts"`
	GetIcons      bool          `json:"getIcons"`
	GetAttributes bool          `json:"getAttributes"`
	GetTariff     bool          `json:"getTariff"`
	LocL          []locationReq `json:"locL"`
}

type locationReq struct {
	Type string `json:"type"`
	Lid  string `json:"lid"`
}

type BVGApiResponse struct {
	SvcResL []SvcRes `json:"svcResL"`
	Err     string   `json:"err,omitempty"`
	ErrTxt  string   `json:"errTxt,omitempty"`
}

type SvcRes struct {
	Meth      string      `json:"meth"`
	Err       string      `json:"err,omitempty"`
	ErrTxt    string      `json:"errTxt,omitempty"`
	ErrTxtOut string      `json:"errTxtOut,omitempty"`
	Res       *SvcResData `json:"res,omitempty"`
}

type SvcResData struct {
	Common CommonData `json:"common"`
	LocL   []Location `json:"locL"`
	JnyL   []Journey  `json:"jnyL,omitempty"`
}

type CommonData struct {
	LocL  []Location   `json:"locL"`
	ProdL []ProdL      `json:"prodL"`
	HimL  []HimMessage `json:"himL,omitempty"`
	RemL  []RemMessage `json:"remL,omitempty"`
}

type RemMessage struct {
	TxtN string `json:"txtN"`
	Code string `json:"code"`
}

type Location struct {
	Lid      string `json:"lid"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	ExtId    string `json:"extId"`
	TZOffset int    `json:"TZOffset,omitempty"`
	ChgTime  string `json:"chgTime,omitempty"`
	State    string `json:"state,omitempty"`
}

type Journey struct {
	Jid     string `json:"jid"`
	Date    string `json:"date"`
	ProdX   int    `json:"prodX"` // index of the product in prodL
	DirTxt  string `json:"dirTxt"`
	DirFlg  string `json:"dirFlg"`
	IsCncl  bool   `json:"isCncl"`
	StbStop struct {
		DPlfS struct {
			Txt  string `json:"txt"`
			Type string `json:"type"`
		} `json:"dPltfS"`
		LocX   int       `json:"locX"` // departure station index
		Idx    int       `json:"idx"`
		DProdX int       `json:"dProdX"`
		DTimeS string    `json:"dTimeS"`           // Planned Departure time
		DTimeR string    `json:"dTimeR,omitempty"` // Prognosed Departure time DTimeS + delay
		MsgL   []MsgList `json:"msgL,omitempty"`   // stop remarks
	} `json:"stbStop"`
	ProdL []struct {
		ProdX int `json:"prodX"`
		FLocX int `json:"fLocX"`
		TLocX int `json:"tLocX"` // index of the stop (direction)
	} `json:"prodL"`
	MsgL []MsgList `json:"msgL,omitempty"` // journey remarks
}

type MsgList struct {
	RemX *int   `json:"remX"`
	HimX *int   `json:"himX"`
	Type string `json:"type"`
}

type ProdL struct {
	Pid     string  `json:"pid"`
	Name    string  `json:"name"`
	NameS   string  `json:"nameS"`
	ProdCtx ProdCtx `json:"prodCtx"`
}

type ProdCtx struct {
	Name    string `json:"name"`
	Line    string `json:"line"`
	LineId  string `json:"lineId,omitempty"`
	CatOutS string `json:"catOutS"`
	CatOut  string `json:"catOut"`
}

type HimMessage struct {
	Head string `json:"head"`
	Text string `json:"text"`
}

func createStationRequest(stationID string) StationRequest {
	return StationRequest{
		BaseRequest: createBaseRequest(),
		SvcReqL: []locDetailsSvcRequest{
			{
				Meth: "LocDetails",
				Req: locDetailsReq{
					GetHIM:        false,
					GetProducts:   true,
					GetIcons:      false,
					GetAttributes: false,
					GetTariff:     false,
					LocL: []locationReq{
						{
							Type: "S",
							Lid:  fmt.Sprintf("@A=1@L=%s@", stationID),
						},
					},
				},
			},
		},
	}
}

func createServerStatusRequest() ServerStatusRequest {
	return ServerStatusRequest{
		BaseRequest: createBaseRequest(),
		SvcReqL: []serverStatusSvcRequest{
			{
				Meth: "ServerInfo",
			},
		},
	}
}

func createBaseRequest() BaseRequest {
	return BaseRequest{
		Lang: "en",
		Ver:  "1.72",
		Ext:  "BVG.1",
		Auth: auth{
			Type: "AID",
			Aid:  "dVg4TZbW8anjx9ztPwe2uk4LVRi9wO",
		},
		Client: client{
			Type: "WEB",
			ID:   "VBB",
			V:    "10003",
			Name: "webapp",
			L:    "vs_webapp",
		},
	}
}

func createDepartureBoardRequest(stationID string, params internal.APIDeparturesParams) DepartureBoardRequest {
	return DepartureBoardRequest{
		BaseRequest: createBaseRequest(),
		SvcReqL: []departureBoardSvcRequest{
			{
				Meth: "StationBoard",
				Req: departureBoardReq{
					StbLoc: locationReq{
						Type: "S",
						Lid:  fmt.Sprintf("@A=1@L=%s@", stationID),
					},
					JnyFltrL: []jnyFltr{
						{
							Type:  "PROD",
							Mode:  "INC",
							Value: 127, // All products
						},
					},
					// Dur: params.Duration,
					Type: "DEP",
					Sort: "PT",
					// MaxJny: params.MaxResultCount,
					// Date:   params.Date,
					// Time:   params.Time,
				},
			},
		},
	}
}
