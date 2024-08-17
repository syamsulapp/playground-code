package main

import (
	"fmt"
	"net/http"
	"time"
)

func ListUrl() {
	links := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://idijakpus.or.id",
		"https://m.idijakpus.or.id",
		"https://dev-simfoni.idijakpus.or.id",
	}

	c := make(chan string)
	for _, link := range links {
		go CheckLink(link, c) // assigned to (go_routine)
	}

	// for i := 0; i < len(links); {
	// 	fmt.Println(<-c) //excec channel base on length of data links
	// }

	//repeating routes
	for l := range c {
		// time.Sleep(5 * time.Second)
		// go CheckLink(l, c)

		go func(link string) { // func literal(function anonymous)
			time.Sleep(5 * time.Second)
			CheckLink(link, c)
		}(l)
	}
}

func CheckLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "maybe site is down")
		c <- "Migth be down url" //channel syntax excec of go_routine
	}

	fmt.Println(link, "is up")
	c <- "yes link is up"
}
