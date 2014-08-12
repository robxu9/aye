package aye

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
