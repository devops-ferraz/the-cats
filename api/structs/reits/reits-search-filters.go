package reits

type SearchFilter struct {
	CompanyID            int     `json:"companyid"`
	CompanyName          string  `json:"companyname"`
	Ticker               string  `json:"ticker"`
	Price                float64 `json:"price"`
	SectorID             int     `json:"sectorid"`
	SectorName           string  `json:"sectorname"`
	SubsectorID          int     `json:"subsectorid"`
	SubsectorName        string  `json:"subsectorname"`
	Segment              string  `json:"segment"`
	SegmentID            int     `json:"segmentid"`
	Gestao               int     `json:"gestao"`
	GestaoF              string  `json:"gestao_f"`
	DY                   float64 `json:"dy"`
	PVP                  float64 `json:"p_vp"`
	ValorPatrimonialCota float64 `json:"valorpatrimonialcota"`
	LiquidezMediaDiaria  float64 `json:"liquidezmediadiaria"`
	PercentualCaixa      float64 `json:"percentualcaixa"`
	DividendCAGR         float64 `json:"dividend_cagr"`
	CotaCAGR             float64 `json:"cota_cagr"`
	NumeroCotistas       float64 `json:"numerocotistas"`
	NumeroCotas          float64 `json:"numerocotas"`
	Patrimonio           float64 `json:"patrimonio"`
	LastDividend         float64 `json:"lastdividend"`
}
