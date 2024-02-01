package bvg

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

type ServerInfoRequest struct {
	baseRequest
	SvcReqL []ServerInfoSvcReq `json:"svcReqL"`
}

func (r *ServerInfoRequest) Create() {
	r.baseRequest = baseRequest{
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
	}
	r.SvcReqL = []ServerInfoSvcReq{
		{
			Meth: "ServerInfo",
			Req:  ServerInfo{GetVersionInfo: "true"},
		},
	}
}

// LocDetailsRequestPayload struct and Create method
type LocDetailsRequestPayload struct {
	baseRequest
	SvcReqL []LocDetailsSvcReq `json:"svcReqL"`
}

func CreateBVGStationRequestPayload(id string) LocDetailsRequestPayload {
	lid := "A=1@L=" + id + "@"
	return LocDetailsRequestPayload{
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
						{Type: "S", Lid: lid},
					},
				},
			},
		},
	}
}

type ServerInfoSvcReq struct {
	Meth string     `json:"meth"`
	Req  ServerInfo `json:"req"`
}

type LocDetailsSvcReq struct {
	Meth string        `json:"meth"`
	Req  LocDetailsReq `json:"req"`
}

type ServerInfo struct {
	GetVersionInfo string `json:"getVersionInfo"`
}

type LocDetailsReq struct {
	LocL []LocationReq `json:"locL"`
}

type LocationReq struct {
	Type string `json:"type"`
	Lid  string `json:"lid"`
}
