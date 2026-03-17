package repo

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	EmptyUser() User

	List() ([]*User, error)
	Store(User) (*User, error)
	Find(email, password string) (*User, error)
}

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

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) EmptyUser() User {
	return User{}
}

func (r *userRepo) List() ([]*User, error) {
	var users []*User
	err := r.db.Select(&users, "SELECT * FROM users ORDER BY id")
	return users, err
}

func (r *userRepo) Store(user User) (*User, error) {
	var id int
	query := `INSERT INTO users (
		first_name,
		last_name,
		email,
		password,
		is_shop_owner
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	) RETURNING id`
	err := r.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.IsShopOwner,
	).Scan(&id)

	if err != nil {
		return nil, err
	}
	user.ID = id

	return &user, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	var user User

	err := r.db.Get(&user, "SELECT * FROM users WHERE email=$1 AND password=$2", email, password)

	if err != nil {
		return nil, err
	}

	return &user, err
}
