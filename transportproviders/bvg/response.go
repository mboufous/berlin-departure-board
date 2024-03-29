package bvg

type ApiResponse struct {
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
	ProdL []ProductReq `json:"prodL"`
	HimL  []HimMessage `json:"himL,omitempty"`
	RemL  []RemMessage `json:"remL,omitempty"`
}

type RemMessage struct {
	TxtN string `json:"txtN"`
}

type Location struct {
	Lid      string `json:"lid"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	ExtId    string `json:"extId"`
	TZOffset int    `json:"TZOffset,omitempty"`
	ChgTime  string `json:"chgTime,omitempty"`
	PCls     int    `json:"pCls,omitempty"`
	PRefL    []int  `json:"pRefL,omitempty"`
	State    string `json:"state,omitempty"`
}

type Journey struct {
	Jid     string `json:"jid"`
	Date    string `json:"date"`
	ProdX   int    `json:"prodX"` // index of the product in prodL
	DirTxt  string `json:"dirTxt"`
	StbStop struct {
		LocX   int    `json:"locX"`
		Idx    int    `json:"idx"` // index of the stop in locL
		DProdX int    `json:"dProdX"`
		DTimeS string `json:"dTimeS"`           // Planned Departure time
		DTimeR string `json:"dTimeR,omitempty"` // Prognosed Departure time delay = DTimeR - DTimeS
	} `json:"stbStop"`
	MsgL  []MsgList `json:"msgL,omitempty"`
	ProdL []struct {
		ProdX int `json:"prodX"`
		FLocX int `json:"fLocX"`
		TLocX int `json:"tLocX"`
		FIdx  int `json:"fIdx"`
		TIdx  int `json:"tIdx"`
	} `json:"prodL"`
	TrainStartDate string `json:"trainStartDate"`
}

type MsgList struct {
	RemX int    `json:"remX"`
	Type string `json:"type"`
}

type ProductReq struct {
	Pid     string     `json:"pid"`
	Name    string     `json:"name"`
	NameS   string     `json:"nameS"`
	ProdCtx ProductCtx `json:"prodCtx"`
	Cls     int        `json:"cls"`
}

type ProductCtx struct {
	Name    string `json:"name"`
	Line    string `json:"line"`
	LineId  string `json:"lineId,omitempty"`
	CatOut  string `json:"catOut"`
	CatOutS string `json:"catOutS"`
	CatOutL string `json:"catOutL"`
}

type HimMessage struct {
	Hid     string `json:"hid"`
	Act     bool   `json:"act"`
	Head    string `json:"head"`
	Text    string `json:"text"`
	Prio    int    `json:"prio"`
	Prod    int    `json:"prod"`
	Src     int    `json:"src"`
	SDate   string `json:"sDate"`
	STime   string `json:"sTime"`
	EDate   string `json:"eDate"`
	ETime   string `json:"eTime"`
	SDaily  string `json:"sDaily"`
	EDaily  string `json:"eDaily"`
	Comp    string `json:"comp"`
	CatRefL []int  `json:"catRefL"`
	PubChL  []struct {
		Name  string `json:"name"`
		FDate string `json:"fDate"`
		FTime string `json:"fTime"`
		TDate string `json:"tDate"`
		TTime string `json:"tTime"`
	} `json:"pubChL"`
}
