package user

type User struct {
	username string
}

func NewUser(username string) *User {
	return &User{username: username}
}
