package main

import (
	"fmt"

	"github.com/go-playground/validator"
)

type User struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Username  string `json:"username" validate:"required,email"`
	Password  string `json:"password" validate:"required,gte=10"`
	Type      string `json:"type,omitempty" validate:"isdefault"`
}

func main() {
	user := User{
		Firstname: "Nic",
		Lastname:  "Raboy",
		Username:  "test@12.com",
		Password:  "1234567890",
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = validate.StructExcept(user, "Firstname", "Lastname")
	if err != nil {
		fmt.Println(err.Error())
	}
	password := "123456w23e"
	// validate = validator.New()
	err = validate.Var(password, "required,gte=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}
