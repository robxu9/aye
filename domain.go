package aye

import (
	"net"

	"fmt"
)

type Domain struct {
	Name     string    `json:"name"`
	TTL      uint64    `json:"ttl"`
	ZoneFile string    `json:"zone_file"`
	client   *DOClient `json:"-"`
}

type Domains struct {
	Domains []*Domain `json:"domains"`
	Meta    *Meta     `json:"meta"`
	Links   *Links    `json:"links,omitempty"`
}

type DomainResult struct {
	Domain *Domain `json:"domain"`
}

func (d *DOClient) CreateDomain(name string, ip net.IP) (*Domain, error) {
	req, err := d.newDefaultRequest("POST", "/v2/domains", map[string]string{
		"name":       name,
		"ip_address": ip.String(),
	})
	if err != nil {
		return nil, err
	}

	result := new(DomainResult)

	_, err = d.DoAndDecode(req, result)

	if err != nil {
		return nil, err
	}

	domain := result.Domain
	domain.client = d

	return domain, nil
}

func (d *DOClient) Domain(name string) (*Domain, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/domains/%s", name), nil)
	if err != nil {
		return nil, err
	}

	result := new(DomainResult)

	_, err = d.DoAndDecode(req, result)

	if err != nil {
		return nil, err
	}

	domain := result.Domain
	domain.client = d

	return domain, nil
}

func (d *Domain) Delete() error {
	req, err := d.client.newDefaultRequest("DELETE", fmt.Sprintf("/v2/domains/%s", d.Name), nil)
	if err != nil {
		return err
	}

	_, err = d.client.Do(req)
	return err
}

func (d *DOClient) ListDomains(page int) (*Domains, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/domains?page=%d", page), nil)
	if err != nil {
		return nil, err
	}

	domains := new(Domains)

	_, err = d.DoAndDecode(req, domains)

	if err != nil {
		return nil, err
	}

	for _, v := range domains.Domains {
		v.client = d
	}

	return domains, nil
}
