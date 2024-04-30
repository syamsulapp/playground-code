package main

import "idiJakartaPusat/routes"

func main() {
	app := routes.Setup()

	app.Listen(":8085")
}
