package main

import (
	"fmt"

	"github.com/kylebustard/lowendcode/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Selim",
		LastName:  "Wormrider",
	}
	fmt.Println(u)
}

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
