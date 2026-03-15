package repo

import "errors"

type UserRepo interface {
	EmptyUser() User

	List() ([]*User, error)
	Store(User) (*User, error)
	Find(email, password string) (*User, error)
}

type User struct {
	ID          int    `json:"id"` // tag
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type userRepo struct {
	userList []*User
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) EmptyUser() User {
	return User{}
}

func (r *userRepo) List() ([]*User, error) {
	return r.userList, nil
}

func (r *userRepo) Store(user User) (*User, error) {
	r.userList = append(r.userList, &user)

	return &user, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	for _, user := range r.userList {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}

	return nil, errors.New("Invalid Creadentails")
}
