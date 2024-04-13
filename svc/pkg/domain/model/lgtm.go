package model

import "time"

type LGTM struct {
	ID        string    `json:"-"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
}
