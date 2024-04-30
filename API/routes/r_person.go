package routes

import (
	"github.com/gofiber/fiber/v2"
)

func indexRoute(c *fiber.Ctx) error {
	return c.SendString("Idi jakarta pusat API")
}

func SetupRoutePerson(app *fiber.App) {
	app.Get("/", indexRoute)
}

func Setup() *fiber.App {
	app := fiber.New()
	SetupRoutePerson(app)
	return app
}
