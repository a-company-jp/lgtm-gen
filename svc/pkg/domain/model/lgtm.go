package model

import "time"

type LGTM struct {
	ID        string    `firestore:"-"`
	Url       string    `firestore:"url"`
	CreatedAt time.Time `firestore:"createdAt"`
}
