package main

import (
	"fmt"
	"regexp"
)

//it works fine
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(e)
}

func main() {
	mail := []string{
		"good@exmaple.com",
		"bad-example@#@.com",
	}
	for _, email := range mail {
		fmt.Printf("%18s valid: %t\n", email, isEmailValid(email))
	}
}
