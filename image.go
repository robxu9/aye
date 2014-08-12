package aye

import "fmt"

type Image struct {
	Id           uint64   `json:"id"`
	Name         string   `json:"name"`
	Distribution string   `json:"distribution"`
	Slug         string   `json:"slug,omitempty"`
	Public       bool     `json:"public"`
	Region       []string `json:"regions"`
	Actions      []uint64 `json:"action_ids"`
	Created      Time     `json:"created_at,omitempty"`

	client *DOClient `json:"-"`
}

type ImageResult struct {
	Image *Image `json:"image"`
}

type Images struct {
	Images []*Image `json:"images"`
	Meta   *Meta    `json:"meta"`
	Links  *Links   `json:"links,omitempty"`
}

func (d *DOClient) Image(id interface{}) (*Image, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/images/%v", id), nil)
	if err != nil {
		return nil, err
	}

	result := new(ImageResult)

	_, err = d.doAndDecode(req, result)
	if err != nil {
		return nil, err
	}

	image := result.Image
	image.client = d
	return image, nil
}

func (i *Image) Delete() error {
	req, err := i.client.newDefaultRequest("DELETE", fmt.Sprintf("/v2/images/%d", i.Id), nil)
	if err != nil {
		return err
	}

	_, err = i.client.do(req)
	return err
}

func (i *Image) Update() error {
	req, err := i.client.newDefaultRequest("PUT", fmt.Sprintf("/v2/images/%d", i.Id), map[string]interface{}{
		"name": i.Name,
	})
	if err != nil {
		return err
	}

	_, err = i.client.do(req)
	return err
}

func (d *DOClient) ListImages(page int) (*Images, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/images?page=%d", page), nil)
	if err != nil {
		return nil, err
	}

	images := new(Images)

	_, err = d.doAndDecode(req, images)
	if err != nil {
		return nil, err
	}

	for _, v := range images.Images {
		v.client = d
	}

	return images, nil
}
