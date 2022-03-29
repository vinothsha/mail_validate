package main

import (
	"fmt"
	"regexp"
)

func main() {
	email := []string{
		"ç$€§/az@gmail.com",
		"abcd@gmail_yahoo.com",
		"abcd@gmail-yahoo.com",
		"abcd@gmailyahoo",
		"abcd@gmail.yahoo",
	}
	for _,mail :=range email {
		Validate(mail)
	}
	
}
func Validate(e string) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	fmt.Printf("Email: %v :%v\n", e, re.MatchString(e))
}
