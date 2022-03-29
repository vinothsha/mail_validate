package main

import (
	"fmt"
	"net/mail"
)

func valid(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
func main() {
	mail:=[]string{
		"good@exmaple.com",
		"bad-example",
	}
	for _, email := range  mail{
		fmt.Printf("%18s valid: %t\n", email, valid(email))
	}

}
