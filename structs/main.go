package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// Based on property order
	// alex := person{"Alex", "Anderson"}

	// Specifying each property
	// alex2 := person{firstName: "Alex", lastName: "Anderson"}

	// Creating a Zero Value instance
	// var alex3 person
	// alex3.lastName = "Alex"
	// alex3.lastName = "Anderson"

	jim := person{
		firstName: "Jim",
		lastName:  "Manao",
		contact: contactInfo{
			email:   "man@mail.com",
			zipCode: 9300,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
