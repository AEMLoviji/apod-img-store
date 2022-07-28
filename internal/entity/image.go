package entity

import "time"

type Image struct {
	ID        string
	Title     string
	Copyright string
	Url       string
	CreatedAt time.Time
}
