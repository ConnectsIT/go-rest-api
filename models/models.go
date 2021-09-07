package models

import "time"

type Img struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Device      string    `json:"device"`
	Color       string    `json:"color"`
	ReleaseDate time.Time `json:"release_date"`
	Runtime     int       `json:"runtime"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
