package models

import (
	"time"
)

type Activity struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Time        time.Time `json:"time"`
	ImagePath   string    `json:"image_path"`
}
