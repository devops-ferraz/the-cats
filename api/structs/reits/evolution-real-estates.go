package reits

type RealEstateFund struct {
	Name           string  `json:"name"`
	CurrentValue   float64 `json:"current_value"`
	Quantity       int     `json:"quantity"`
	Dividends      float64 `json:"dividends"`
	Months         int     `json:"months"`
	MonthlyDeposit float64 `json:"monthly_deposit"`
}

type Evolution struct {
	Month          int     `json:"month"`
	QuantityShares int     `json:"quantity_shares"`
	Dividends      float64 `json:"dividends"`
	CurrentYield   float64 `json:"current_yield"`
	Change         float64 `json:"change"`
	ShareValue     float64 `json:"share_value"`
}
