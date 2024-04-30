package main

import "fmt"

func main() {

	// variable standart
	var nama string = "samsul"
	var nim = 1

	kuliah := "marif"

	fmt.Println(nama + kuliah)
	fmt.Println(nim)
	fmt.Println(kuliah)

	// multiple variable
	var a, b, c int = 2, 4, 6

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	var (
		x int = 2
		y int = 3
	)

	fmt.Println(x + y)

	//naming rules
	var myCampusTheBest string = "universitasHaluoleo"
	var my_favorite_programming string = "go php js"
	var MyFriendsJob string = "majid"

	fmt.Println(myCampusTheBest)
	fmt.Println(my_favorite_programming)
	fmt.Println(MyFriendsJob)

}
