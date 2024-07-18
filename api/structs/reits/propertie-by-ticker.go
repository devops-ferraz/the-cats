package reits

type TenantSector struct {
	Type                      string  `json:"type"`
	PercentPropertyRevenue    float64 `json:"percentPropertyRevenue"`
	PercentTenantToFiiRevenue float64 `json:"percentTenantToFiiRevenue"`
}

type Property struct {
	Type                   string         `json:"type"`
	Name                   string         `json:"name"`
	DatePurchase           *string        `json:"datePurchase,omitempty"`
	PercentagePartic       float64        `json:"percentagePartic"`
	ValueGrossLeasableArea float64        `json:"valueGrossLeasableArea"`
	State                  string         `json:"state"`
	City                   string         `json:"city"`
	Address                string         `json:"address"`
	GoogleMapsLink         string         `json:"googleMapsLink"`
	PercentVacancy         float64        `json:"percentVacancy"`
	Percent90DayDeliquency float64        `json:"percent90DayDeliquency"`
	PercentFii             float64        `json:"percentFii"`
	TenantSector           []TenantSector `json:"tenantSector"`
}

type ReitProperty struct {
	Ticker   string     `json:"ticker"`
	Property []Property `json:"property"`
}
