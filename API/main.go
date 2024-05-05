package main

import "idiJakartaPusat/Routes"

func main() {
	app := Routes.Setup()

	app.Listen(":8085")
}
