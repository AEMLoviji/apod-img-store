package entity

import "time"

type Image struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Copyright string    `json:"copyright"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
}
