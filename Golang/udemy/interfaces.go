package main

import (
	"fmt"
	"net/http"
)

type Bot interface {
	GetGreeting() string
}
type GetUrl interface {
	GetUrlTest() string
}
type EnglishBot struct{}
type SpanishBot struct{}
type HitUrl struct{}

func PrintGreeting(b Bot) {
	fmt.Println(b.GetGreeting())
}

func PrintUrl(HandlerUrl GetUrl) {
	// fmt.Println(http.Get(HandlerUrl.GetUrlTest()))
	res, err := http.Get(HandlerUrl.GetUrlTest())
	if err != nil {
		fmt.Println("Error", err)
	}

	bs := make([]byte, 99999)
	res.Body.Read(bs)
	fmt.Println(string(bs))
}

func (EnglishBot) GetGreeting() string {
	return "hi there!"
}

func (sb SpanishBot) GetGreeting() string {
	return "hola!"
}

func (HandlerUrl HitUrl) GetUrlTest() string {
	return "https://www.google.com"
}
