package models

import (
	"log"
)

type User struct {
	Id int
	Username string
}

func FindUser(number int) (bool, *User) {
	obj, err := dbmap.Get(User{}, number)
	emp := obj.(*User)

	if err != nil {
		log.Print("ERROR findUser: ")
		log.Println(err)
	}

	return (err == nil), emp
}

func CreateUser(username string) (bool) {

	user := &User{0,username}
	dbmap.Insert(user)

	return true
}