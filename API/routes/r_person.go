package routes

import (
	"idiJakartaPusat/controller"

	"github.com/gofiber/fiber/v2"
)

func indexRoute(c *fiber.Ctx) error {
	return c.SendString("Idi jakarta pusat API")
}

func SetupRoutePerson(app *fiber.App) {
	app.Get("/", indexRoute)
	app.Get("/person", controller.GetPersonWithObjectStatic)
}

func Setup() *fiber.App {
	app := fiber.New()
	SetupRoutePerson(app)
	return app
}
