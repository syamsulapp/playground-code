package main

import (
	"fmt"
	"io"
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

type LogUrl struct{}

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
	// fmt.Println(http.Get(HandlerUrl.Url))
	res, err := http.Get(HandlerUrl.Url)
	if err != nil {
		fmt.Println("Error", err)
	}

	lu := LogUrl{}
	io.Copy(lu, res.Body)
}

func (LogUrl) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
