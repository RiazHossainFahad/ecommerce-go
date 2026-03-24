package repo

import (
	"ecommerce/domain"
	"ecommerce/user"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) EmptyUser() domain.User {
	return domain.User{}
}

func (r *userRepo) List() ([]*domain.User, error) {
	var users []*domain.User
	err := r.db.Select(&users, "SELECT * FROM users ORDER BY id")
	return users, err
}

func (r *userRepo) Store(user domain.User) (*domain.User, error) {
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

func (r *userRepo) Find(email, password string) (*domain.User, error) {
	var user domain.User

	err := r.db.Get(&user, "SELECT * FROM users WHERE email=$1 AND password=$2", email, password)

	if err != nil {
		return nil, err
	}

	return &user, err
}
