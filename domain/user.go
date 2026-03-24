package domain

import "time"

type User struct {
	ID          int        `json:"id" db:"id"` // tag
	FirstName   string     `json:"first_name" db:"first_name"`
	LastName    string     `json:"last_name" db:"last_name"`
	Email       string     `json:"email" db:"email"`
	Password    string     `json:"password" db:"password"`
	IsShopOwner bool       `json:"is_shop_owner" db:"is_shop_owner"`
	CreatedAt   *time.Time `json:"-" db:"created_at"`
	UpdatedAt   *time.Time `json:"-" db:"updated_at"`
}
