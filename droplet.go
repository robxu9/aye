package aye

import "fmt"

type Kernel struct {
	Id      uint64 `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type Droplet struct {
	Id          uint64  `json:"id"`
	Name        string  `json:"name"`
	Memory      uint64  `json:"memory"`
	VirtualCPUs uint8   `json":"vcpus"`
	Disk        uint64  `json:"disk"`
	Region      *Region `json:"region"`
	Image       *Image  `json:"image"`
	Kernel      *Kernel `json:"kernel"`
	Size        *Size   `json:"size"`
	Locked      bool    `json:"locked"`
	Created     Time    `json:"created_at"`
	Status      string  `json:"status"`
	Networks    struct {
		V4 []struct {
			IP      string `json:"ip_address"`
			Netmask string `json:"netmask"`
			Gateway string `json:"gateway"`
			Type    string `json:"type"`
		} `json:"v4"`
		V6 []struct {
			IP      string `json:"ip_address"`
			CIDR    int    `json:"cidr"`
			Gateway string `json:"gateway"`
			Type    string `json:"type"`
		}
	} `json:"networks"`
	Backups   []uint64 `json:"backup_ids"`
	Snapshots []uint64 `json:"snapshot_ids"`
	Actions   []uint64 `json:"action_ids"`
	Features  []string `json:"features"`

	client *DOClient `json:"-"`
}

type DropletCreate struct {
	Name       string      `json:"name"`
	Region     string      `json:"region"`
	Size       string      `json:"size"`
	Image      interface{} `json:"image"`
	SSHKeys    []string    `json:"ssh_keys,omitempty"`
	Backup     bool        `json:"backups,omitempty"`
	IPv6       bool        `json:"ipv6,omitempty"`
	PrivateNet bool        `json:"private_networking,omitempty"`
}

type Droplets struct {
	Droplets []*Droplet `json:"droplets"`
	Meta     *Meta      `json:"meta"`
	Links    *Links     `json:"links,omitempty"`
}

type DropletResult struct {
	Droplet *Droplet `json:"droplet"`
}

type Kernels struct {
	Kernels []*Kernel `json:"kernels"`
	Meta    *Meta     `json:"meta"`
	Links   *Links    `json:"links,omitempty"`
}

type Snapshots struct {
	Snapshots []*Image `json:"snapshots"`
	Meta      *Meta    `json:"meta"`
	Links     *Links   `json:"links,omitempty"`
}

type Backups struct {
	Backups []*Image `json:"backups"`
	Meta    *Meta    `json:"meta"`
	Links   *Links   `json:"links,omitempty"`
}

func (d *DOClient) ListDroplets(page int) (*Droplets, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/droplets?page=%d", page), nil)
	if err != nil {
		return nil, err
	}

	droplets := new(Droplets)

	_, err = d.doAndDecode(req, droplets)

	if err != nil {
		return nil, err
	}

	for _, v := range droplets.Droplets {
		v.client = d
	}

	return droplets, nil
}

func (d *DOClient) CreateDroplet(params *DropletCreate) (*Droplet, error) {
	req, err := d.newDefaultRequest("POST", "/v2/droplets", params)
	if err != nil {
		return nil, err
	}

	result := new(DropletResult)

	_, err = d.doAndDecode(req, result)
	if err != nil {
		return nil, err
	}

	droplet := result.Droplet
	droplet.client = d
	return droplet, nil
}

func (d *DOClient) Droplet(id uint64) (*Droplet, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/droplets/%d", id), nil)
	if err != nil {
		return nil, err
	}

	result := new(DropletResult)

	_, err = d.doAndDecode(req, result)
	if err != nil {
		return nil, err
	}

	droplet := result.Droplet
	droplet.client = d
	return droplet, nil
}

func (d *Droplet) Delete() error {
	req, err := d.client.newDefaultRequest("DELETE", fmt.Sprintf("/v2/droplets/%d", d.Id), nil)
	if err != nil {
		return err
	}

	_, err = d.client.do(req)
	return err
}

func (d *Droplet) ListKernels(page int) (*Kernels, error) {
	req, err := d.client.newDefaultRequest("GET", fmt.Sprintf("/v2/droplets/%d/kernels?page=%d", d.Id, page), nil)
	if err != nil {
		return nil, err
	}

	kernels := new(Kernels)

	_, err = d.client.doAndDecode(req, kernels)

	if err != nil {
		return nil, err
	}

	return kernels, nil
}

func (d *Droplet) ListSnapshots(page int) (*Snapshots, error) {
	req, err := d.client.newDefaultRequest("GET", fmt.Sprintf("/v2/droplets/%d/snapshots?page=%d", d.Id, page), nil)
	if err != nil {
		return nil, err
	}

	snapshots := new(Snapshots)

	_, err = d.client.doAndDecode(req, snapshots)

	if err != nil {
		return nil, err
	}

	for _, v := range snapshots.Snapshots {
		v.client = d.client
	}

	return snapshots, nil
}

func (d *Droplet) ListBackups(page int) (*Backups, error) {
	req, err := d.client.newDefaultRequest("GET", fmt.Sprintf("/v2/droplets/%d/backups?page=%d", d.Id, page), nil)
	if err != nil {
		return nil, err
	}

	backups := new(Backups)

	_, err = d.client.doAndDecode(req, backups)

	if err != nil {
		return nil, err
	}

	for _, v := range backups.Backups {
		v.client = d.client
	}

	return backups, nil
}

func (d *Droplet) ListActions(page int) (*Actions, error) {
	req, err := d.client.newDefaultRequest("GET", fmt.Sprintf("/v2/droplets/%d/actions?page=%d", d.Id, page), nil)
	if err != nil {
		return nil, err
	}

	actions := new(Actions)

	_, err = d.client.doAndDecode(req, actions)

	if err != nil {
		return nil, err
	}

	return actions, nil
}
