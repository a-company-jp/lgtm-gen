package model

import "time"

type LGTM struct {
	ID        string
	CreatedAt time.Time `firestore:"createdAt"`
}
