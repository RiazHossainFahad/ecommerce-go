package db

type User struct {
	ID          int    `json:"id"` // tag
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

var userList []User

func GetUserList() []User {
	return userList
}

func StoreUser(user User) {
	userList = append(userList, user)
}

func FindUser(email, password string) *User {
	for _, user := range userList {
		if user.Email == email && user.Password == password {
			return &user
		}
	}

	return nil
}
