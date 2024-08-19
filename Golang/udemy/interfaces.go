package main

import (
	"fmt"
	"io"
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
type HitUrl struct {
	Url string
}

func PrintGreeting(b Bot) {
	fmt.Println(b.GetGreeting())
}

func PrintUrl(HandlerUrl GetUrl) {
	// fmt.Println(http.Get(HandlerUrl.GetUrlTest()))
	res, err := http.Get(HandlerUrl.GetUrlTest())
	if err != nil {
		fmt.Println("Error", err)
	}

	lu := LogUrl{}
	io.Copy(lu, res.Body)
}

func (EnglishBot) GetGreeting() string {
	return "hi there!"
}

func (sb SpanishBot) GetGreeting() string {
	return "hola!"
}

func (HandlerUrl HitUrl) GetUrlTest() string {
	return HandlerUrl.Url
}

func (HandlerUrl *HitUrl) UpdateUrl(NewUrl string) {
	(*HandlerUrl).Url = NewUrl
}
