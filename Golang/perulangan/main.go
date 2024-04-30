package main

import "fmt"

func main() {

	for x := 0; x <= 5; x++ {
		if x == 4 {
			continue
		}
		fmt.Println(x)
	}

	for x := 5; x >= 0; x-- {
		if x == 2 {
			break
		}
		fmt.Println(x)
	}

}
