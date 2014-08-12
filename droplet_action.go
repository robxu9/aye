package aye

import "fmt"

func (d *Droplet) doAction(t string, params map[string]interface{}) (*Action, error) {
	if params == nil {
		params = make(map[string]interface{})
	}
	params["type"] = t

	req, err := d.client.newDefaultRequest("POST", fmt.Sprintf("/v2/droplets/%s/actions", d.Id), params)
	if err != nil {
		return nil, err
	}

	action := new(ActionResult)

	_, err = d.client.DoAndDecode(req, action)
	if err != nil {
		return nil, err
	}

	return action.Action, nil
}

func (d *Droplet) Reboot() (*Action, error) {
	return d.doAction("reboot", nil)
}

func (d *Droplet) PowerCycle() (*Action, error) {
	return d.doAction("power_cycle", nil)
}

func (d *Droplet) Shutdown() (*Action, error) {
	return d.doAction("shutdown", nil)
}

func (d *Droplet) PowerOff() (*Action, error) {
	return d.doAction("power_off", nil)
}

func (d *Droplet) PowerOn() (*Action, error) {
	return d.doAction("power_on", nil)
}

func (d *Droplet) PasswordReset() (*Action, error) {
	return d.doAction("password_reset", nil)
}

func (d *Droplet) Resize(size string) (*Action, error) {
	return d.doAction("resize", map[string]interface{}{
		"size": size,
	})
}

func (d *Droplet) Restore(image uint64) (*Action, error) {
	return d.doAction("restore", map[string]interface{}{
		"image": image,
	})
}

func (d *Droplet) Rebuild(image interface{}) (*Action, error) {
	return d.doAction("restore", map[string]interface{}{
		"image": image,
	})
}

func (d *Droplet) Rename(name string) (*Action, error) {
	return d.doAction("rename", map[string]interface{}{
		"name": name,
	})
}

func (d *Droplet) ChangeKernel(kernel uint64) (*Action, error) {
	return d.doAction("change_kernel", map[string]interface{}{
		"kernel": kernel,
	})
}

func (d *Droplet) EnableIPv6() (*Action, error) {
	return d.doAction("enable_ipv6", nil)
}

func (d *Droplet) DisableBackups() (*Action, error) {
	return d.doAction("disable_backups", nil)
}

func (d *Droplet) EnablePrivateNet() (*Action, error) {
	return d.doAction("enable_private_networking", nil)
}

func (d *Droplet) Snapshot(name string) (*Action, error) {
	return d.doAction("snapshot", map[string]interface{}{
		"name": name,
	})
}

func (d *Droplet) Action(id uint64) (*Action, error) {
	req, err := d.client.newDefaultRequest("GET", fmt.Sprintf("/v2/droplets/%d/actions/%d", d.Id, id), nil)
	if err != nil {
		return nil, err
	}

	action := new(ActionResult)

	_, err = d.client.DoAndDecode(req, action)
	if err != nil {
		return nil, err
	}

	return action.Action, nil
}
