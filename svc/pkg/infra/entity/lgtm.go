package entity

import "time"

type LGTM struct {
	ID        string
	Title     string    `firestore:"title"`
	CreatedAt time.Time `firestore:"createdAt"`
}
