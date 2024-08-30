package models

import (
	"time"
)

type Event struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Time        time.Time `json:"time"`
	ImageURL    string    `json:"image_url"`
}
