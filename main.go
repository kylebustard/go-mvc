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
