package main

import (
	"fmt"
	"net/http"
)

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

type StoreUrl struct {
	Url string
}

func (PointerToPerson *Person) ChangeName(NewFirstName string) {
	(*PointerToPerson).FirstName = NewFirstName
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func (PointerChangingUrl *StoreUrl) ChangeUrl(NewUrl string) {
	(*PointerChangingUrl).Url = NewUrl
}

func (HandlerUrl StoreUrl) PrintResponseUrl() {
	fmt.Println(http.Get(HandlerUrl.Url))
}
