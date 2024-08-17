package main

import (
	"fmt"
	"net/http"
)

func ListUrl() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://idijakpus.or.id",
		"https://m.idijakpus.or.id",
		"https://dev-simfoni.idijakpus.or.id",
	}

	for _, link := range links {
		CheckLink(link)
	}
}

func CheckLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "maybe site is down")
	}

	fmt.Println(link, "is up")
}
