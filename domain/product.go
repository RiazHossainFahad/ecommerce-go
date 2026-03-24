package domain

import "time"

type Product struct {
	ID          int        `json:"id" db:"id"` // tag
	Title       string     `json:"title" db:"title"`
	Description string     `json:"description" db:"description"`
	Price       float64    `json:"price" db:"price"`
	ImgUrl      string     `json:"img_url" db:"img_url"`
	CreatedAt   *time.Time `json:"-" db:"created_at"`
	UpdatedAt   *time.Time `json:"-" db:"updated_at"`
}
