package main

import "fmt"

func main() {
	x := 2
	y := 2

	if (x + y) == 5 {
		fmt.Println("hasil benar")
	} else {
		fmt.Println("hasil salah")
	}

	if (x + y) == 6 {
		fmt.Println("hasil salah")
	} else if (x + y) == 5 {
		fmt.Println("hasil benar")
	} else {
		fmt.Println("hasil jauh")
	}

	if (x+y) == 5 && (x+y) >= 4 {
		if (x + 5) != 6 {
			fmt.Println("hasil benar")
		}
	} else {
		fmt.Println("hasil salah")
	}
}
