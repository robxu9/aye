package aye

import "fmt"

type Key struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
	PublicKey   string `json:"public_key"`

	client *DOClient `json:"-"`
}

type Keys struct {
	Keys  []*Key `json:"ssh_keys"`
	Meta  *Meta  `json:"meta"`
	Links *Links `json:"links,omitempty"`
}

type KeyResult struct {
	Key *Key `json:"ssh_key"`
}

func (d *DOClient) CreateKey(name, publickey string) (*Key, error) {
	req, err := d.newDefaultRequest("GET", "/v2/account/keys", map[string]interface{}{
		"name":       name,
		"public_key": publickey,
	})

	result := new(KeyResult)

	_, err = d.doAndDecode(req, result)
	if err != nil {
		return nil, err
	}

	key := result.Key
	key.client = d
	return key, nil
}

func (k *Key) Delete() error {
	req, err := k.client.newDefaultRequest("DELETE", fmt.Sprintf("/v2/account/keys/%d", k.Id), nil)
	if err != nil {
		return err
	}

	_, err = k.client.do(req)
	return err
}

func (d *DOClient) Key(id_or_fingerprint interface{}) (*Key, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/account/keys/%v", id_or_fingerprint), nil)

	result := new(KeyResult)

	_, err = d.doAndDecode(req, result)
	if err != nil {
		return nil, err
	}

	key := result.Key
	key.client = d
	return key, nil
}

func (k *Key) Update() error {
	req, err := k.client.newDefaultRequest("PUT", fmt.Sprintf("/v2/account/keys/%d", k.Id), map[string]interface{}{
		"name": k.Name,
	})
	if err != nil {
		return err
	}

	_, err = k.client.do(req)
	return err
}

func (d *DOClient) ListKeys(page int) (*Keys, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/account/keys?page=%d", page), nil)
	if err != nil {
		return nil, err
	}

	keys := new(Keys)

	_, err = d.doAndDecode(req, keys)
	if err != nil {
		return nil, err
	}

	for _, v := range keys.Keys {
		v.client = d
	}

	return keys, nil
}
