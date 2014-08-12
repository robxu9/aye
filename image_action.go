package aye

import "fmt"

func (i *Image) doAction(t string, params map[string]interface{}) (*Action, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	params["type"] = t

	req, err := i.client.newDefaultRequest("POST", fmt.Sprintf("/v2/images/%s/actions", i.Id), params)
	if err != nil {
		return nil, err
	}

	action := new(ActionResult)

	_, err = i.client.doAndDecode(req, action)
	if err != nil {
		return nil, err
	}

	return action.Action, nil
}

func (i *Image) Transfer(region string) (*Action, error) {
	return i.doAction("transfer", map[string]interface{}{
		"region": region,
	})
}

func (i *Image) Action(id uint64) (*Action, error) {
	req, err := i.client.newDefaultRequest("GET", fmt.Sprintf("/v2/images/%d/actions/%d", i.Id, id), nil)
	if err != nil {
		return nil, err
	}

	action := new(ActionResult)

	_, err = i.client.doAndDecode(req, action)
	if err != nil {
		return nil, err
	}

	return action.Action, nil
}
