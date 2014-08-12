package aye

import "fmt"

type Action struct {
	Id           uint64 `json:"id"`
	Status       string `json:"status"`
	Type         string `json:"type"`
	Started      Time   `json:"started_at"`
	Completed    Time   `json:"completed_at"`
	ResourceId   uint64 `json:"resource_id"`
	ResourceType string `json:"resource_type"`
	Region       string `json:"region"`
}

type Actions struct {
	Actions []*Action `json:"actions"`
	Meta    *Meta     `json:"meta"`
	Links   *Links    `json:"links,omitempty"`
}

type ActionResult struct {
	Action *Action `json:"action"`
}

func (d *DOClient) Action(id int) (*Action, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/actions/%d", id), nil)
	if err != nil {
		return nil, err
	}

	action := new(ActionResult)

	_, err = d.DoAndDecode(req, action)

	if err != nil {
		return nil, err
	}

	return action.Action, nil
}

func (d *DOClient) ListActions(page int) (*Actions, error) {
	req, err := d.newDefaultRequest("GET", fmt.Sprintf("/v2/actions?page=%d", page), nil)
	if err != nil {
		return nil, err
	}

	actions := new(Actions)

	_, err = d.DoAndDecode(req, actions)

	if err != nil {
		return nil, err
	}

	return actions, nil
}
