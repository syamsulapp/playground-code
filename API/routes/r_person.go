package routes

import (
	"idiJakartaPusat/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

func indexRoute(c *fiber.Ctx) error {
	return c.SendString("IDI Jakarta Pusat")
}

func SetupRoutePerson(app *fiber.App) {
	app.Get("/", indexRoute)
	app.Get("/person", controller.GetPersonWithObjectStatic)
}

func Setup() *fiber.App {
	fastergoding.Run()
	app := fiber.New()
	SetupRoutePerson(app)
	return app
}
