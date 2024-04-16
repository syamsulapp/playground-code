package main

import "fmt"

func main() {
	x := 2
	y := 2

	hasil := x + y

	switch hasil {
	case 5:
		fmt.Println("benar")
	default:
		fmt.Println("hasil salah")
	}
}
