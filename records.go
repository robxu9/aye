package aye

import "fmt"

type DomainRecord struct {
	Id       uint64 `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	Priority uint16 `json:"priority,omitempty"`
	Port     uint16 `json:"port,omitempty"`
	Weight   uint16 `json:"weight,omitempty"`

	client *DOClient `json:"-"`
	domain string    `json:"-"`
}

type DomainRecords struct {
	Records []*DomainRecord `json:"domain_records"`
	Meta    *Meta           `json:"meta"`
	Links   *Links          `json:"links,omitempty"`
}

type DomainRecordResult struct {
	Record *DomainRecord `json:"domain_record"`
}

func (d *DOClient) CreateDomainRecord(domain string, record *DomainRecord) (*DomainRecord, error) {
	req, err := d.newDefaultRequest("POST", fmt.Sprintf("/v2/domains/%s/records", domain), record)
	if err != nil {
		return nil, err
	}

	result := new(DomainRecordResult)

	_, err = d.DoAndDecode(req, result)

	if err != nil {
		return nil, err
	}

	r := result.Record
	r.client = d
	r.domain = domain

	return r, nil
}

func (d *DomainRecord) Delete() error {
	req, err := d.client.newDefaultRequest("DELETE", fmt.Sprintf("/v2/domains/%s/records/%d", d.domain, d.Id), nil)
	if err != nil {
		return err
	}

	_, err = d.client.Do(req)

	return err
}

func (d *DOClient) DomainRecord(domain string, id uint64) (*DomainRecord, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/domains/%s/records/%d", domain, id), nil)
	if err != nil {
		return nil, err
	}

	result := new(DomainRecordResult)

	_, err = d.DoAndDecode(req, result)

	if err != nil {
		return nil, err
	}

	r := result.Record
	r.client = d
	r.domain = domain

	return r, nil
}

func (d *DomainRecord) Update() error {
	req, err := d.client.newDefaultRequest("PUT", fmt.Sprintf("/v2/domains/%s/records/%d", d.domain, d.Id), map[string]string{
		"name": d.Name,
	})

	if err != nil {
		return err
	}

	_, err = d.client.Do(req)
	return err
}

func (d *DOClient) ListDomainRecords(domain string, page int) (*DomainRecords, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/domains/%s/records?page=%d", domain, page), nil)
	if err != nil {
		return nil, err
	}

	records := new(DomainRecords)

	_, err = d.DoAndDecode(req, records)

	if err != nil {
		return nil, err
	}

	for _, v := range records.Records {
		v.client = d
		v.domain = domain
	}

	return records, nil
}
