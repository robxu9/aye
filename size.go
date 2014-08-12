package aye

import "fmt"

type Size struct {
	Slug         string   `json:"slug"`
	Memory       uint64   `json:"memory"`
	VirtualCPUs  uint64   `json:"vcpus"`
	Disk         uint64   `json:"disk"`
	Transfer     uint64   `json:"transfer"`
	MonthlyPrice float32  `json:"price_monthly"`
	HourlyPrice  float32  `json:"price_hourly"`
	Regions      []string `json:"regions"`
}

type Sizes struct {
	Sizes []*Size `json:"sizes"`
	Meta  *Meta   `json:"meta"`
	Links *Links  `json:"links,omitempty"`
}

func (d *DOClient) ListSizes(page int) (*Sizes, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/sizes?page=%d", page), nil)
	if err != nil {
		return nil, err
	}

	sizes := new(Sizes)
	_, err = d.doAndDecode(req, sizes)
	if err != nil {
		return nil, err
	}

	return sizes, nil
}
