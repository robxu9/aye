package aye

import "fmt"

type Region struct {
	Slug      string   `json:"slug"`
	Name      string   `json:"name"`
	Sizes     []string `json:"sizes"`
	Available bool     `json:"available"`
	Features  []string `json:"features"`
}

type Regions struct {
	Regions []*Region `json:"regions"`
	Meta    *Meta     `json:"meta"`
	Links   *Links    `json:"links,omitempty"`
}

func (d *DOClient) ListRegions(page int) (*Regions, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/regions?page=%d", page), nil)
	if err != nil {
		return nil, err
	}

	regions := new(Regions)
	_, err = d.doAndDecode(req, regions)
	if err != nil {
		return nil, err
	}

	return regions, nil
}
