package model

import "time"

type LGTM struct {
	ID        string    `firestore:"-"`
	CreatedAt time.Time `firestore:"createdAt"`
}
