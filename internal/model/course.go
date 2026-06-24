package model

import "time"

type Course struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Price    int    `json:"price"`
	IsActive bool   `json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCourse struct {
	Title    string `json:"title" binding:"required"`
	Price    int    `json:"price"`
	IsActive bool   `json:"is_active"`
}

type UpdateCourse struct {
	Title    *string `json:"title" binding:"required"`
	Price    *int    `json:"price"`
	IsActive *bool   `json:"is_active"`
}
