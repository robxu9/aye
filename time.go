package aye

import "time"

const (
	DOTIME = "2006-01-02T15:04:05Z"
)

type Time struct {
	time.Time
}

func (j Time) format() string {
	return j.Time.Format(DOTIME)
}

func (j Time) MarshalText() ([]byte, error) {
	return []byte(j.format()), nil
}

func (j Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + j.format() + `"`), nil
}

func (j Time) String() string {
	return j.String()
}
