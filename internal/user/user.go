package user

import (
	"encoding/json"
	"io"
	"os"
)

const path = "internal/user/usersdb.json"

type User struct {
	Username string
}

func ReadUsers() (users []User, err error) {
	var f *os.File

	if f, err = os.OpenFile(path, os.O_RDONLY, 0644); err != nil {
		return
	}

	var bytes []byte
	if bytes, err = io.ReadAll(f); err != nil {
		return
	}

	if err = json.Unmarshal(bytes, &users); err != nil {
		return users, err
	}

	return
}

func WriteUsers(users User) (err error) {
	var f *os.File

	if isUserInDb(users) {
		return
	}

	if f, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644); err != nil {
		return
	}
	usersArray, err := ReadUsers()
	usersArray = append(usersArray, users)
	var bytes []byte

	if bytes, err = json.Marshal(usersArray); err != nil {
		return
	}

	if _, err = f.Write(bytes); err != nil {
		return
	}

	return
}

func isUserInDb(user User) bool {
	Users, err := ReadUsers()
	if err != nil {
		panic(err)
	}
	for i := range Users {
		if Users[i].Username == user.Username {
			return true
		}
	}
	return false
}
