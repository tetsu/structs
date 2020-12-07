package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var alex person

	alex = person{
		firstName: "Alex",
		lastName:  "Anderson",
		contact: contactInfo{
			email: "email.com",
			zip:   12345,
		},
	}

	fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
