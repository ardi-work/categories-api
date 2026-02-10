package models

import "time"

type Category struct {
	ID          int       `json:"id" example:"1"`
	Name        string    `json:"name" example:"Makanan"`
	Description string    `json:"description" example:"Kategori makanan"`
	CreatedAt   time.Time `json:"created_at" example:"2026-02-10T10:00:00Z"`
	UpdatedAt   time.Time `json:"updated_at" example:"2026-02-10T10:00:00Z"`
}
