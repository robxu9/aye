package aye

type Size struct {
	Slug         string   `json:"slug"`
	Transfer     uint64   `json:"transfer"`
	MonthlyPrice float32  `json:"price_monthly"`
	HourlyPrice  float32  `json:"price_hourly"`
	Regions      []string `json:"regions"`
}
