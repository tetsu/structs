package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	var alex person

	alex = person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email: "email.com",
			zip:   12345,
		},
	}

	alex.print()

	alexPointer := &alex
	alexPointer.updateName("Alexander")
	alex.print()
}

func (p person) print (){
	fmt.Printf("%+v", p)
	fmt.Println("")
}

func (pointerToPerson *person) updateName (newFirstName string){
	(*pointerToPerson).firstName = newFirstName
}