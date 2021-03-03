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
	// dave := person{"Dave", "Hernandez"}
	// var dave person
	// dave.firstName = "DAVE"
	// dave.lastName = "The best"
	// dave.contact.email = "dave@yes.com"
	// dave.contact.zipCode = 170608

	dave := person{
		firstName: "DAVE",
		lastName:  "THE BEST",
		contact: contactInfo{
			email:   "k@o.com",
			zipCode: 170605,
		},
	}
	dave.updateName("edison")
	dave.print()
}

func (pointerToPerson *person) updateName(newFR string) {
	(*pointerToPerson).firstName = newFR
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
