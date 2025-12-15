package models
type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	CreatedAt string `json:"created_at"`
}

var Users []User

func (u *User) Save() {
	Users = append(Users, *u)
}

func GetAllUsers() []User {
	return Users
}