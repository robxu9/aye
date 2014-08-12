package aye

import (
	"bytes"
	"encoding/json"
	"errors"

	"net/http"
)

var (
	ErrRateLimit    = errors.New("aye: exceeded rate limit")
	ErrNotFound     = errors.New("aye: not found")
	ErrUnauthorized = errors.New("aye: couldn't authenticate")
	ErrTheirEnd     = errors.New("aye: their end has a problem")
	ErrUnknown      = errors.New("aye: unknown status response code")
)

type DOClient struct {
	Token  string
	Client *http.Client
}

func (d *DOClient) newDefaultRequest(method, path string, v interface{}) (*http.Request, error) {
	return d.newRequest(method, "https://api.digitalocean.com/"+path, v)
}

func (d *DOClient) newRequest(method, urlStr string, v interface{}) (*http.Request, error) {
	if d.Client == nil {
		d.Client = http.DefaultClient
	}

	body := &bytes.Buffer{}

	if v != nil {
		encoder := json.NewEncoder(body)

		err := encoder.Encode(v)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+d.Token)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (d *DOClient) do(req *http.Request) (*http.Response, error) {
	resp, err := d.Client.Do(req)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode {
	case http.StatusOK, http.StatusNoContent:
		return resp, nil
	case http.StatusNotFound:
		return nil, ErrNotFound
	case 429:
		return nil, ErrRateLimit
	case http.StatusUnauthorized:
		return nil, ErrUnauthorized
	case http.StatusBadGateway, http.StatusInternalServerError:
		return nil, ErrTheirEnd
	default:
		return nil, ErrUnknown
	}
}

func (d *DOClient) doAndDecode(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := d.do(req)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(resp.Body)

	err = decoder.Decode(v)
	if err != nil {
		return nil, err
	}

	return resp, err
}
