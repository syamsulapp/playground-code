package main

import "fmt"

// define struct
type Person struct {
	FirstName string
	LastName  string
	ContactInfo
}

type ContactInfo struct {
	Email   string
	ZipCode int
}

func (PointerToPerson *Person) ChangeName(NewFirstName string) {
	(*PointerToPerson).FirstName = NewFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}
