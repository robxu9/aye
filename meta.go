package aye

type Meta struct {
	Total uint64 `json:"total"`
}

type Links struct {
	Pages struct {
		First string `json:"first"`
		Prev  string `json:"prev"`
		Next  string `json:"next"`
		Last  string `json:"last"`
	} `json:"pages,omitempty"`
	Actions struct {
		Id   uint64 `json:"id"`
		Type string `json:"rel"`
		Url  string `json:"href"`
	} `json:"actions,omitempty"`
}
