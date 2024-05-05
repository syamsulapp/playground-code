package Routes

import (
	"idiJakartaPusat/Controller"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

func indexRoute(c *fiber.Ctx) error {
	return c.SendString("IDI API Jakarta Pusat")
}

func SetupRoutePerson(app *fiber.App) {
	app.Get("/", indexRoute)

	v1 := app.Group("/api/v1")
	v1.Get("/person/static", Controller.PersonWithObjectStatic)
	v1.Get("/person/body", Controller.PersonWithObjectBody)
	v1.Get("/person/query", Controller.PersonWithQueryParams)
}

func Setup() *fiber.App {
	fastergoding.Run()
	app := fiber.New()
	SetupRoutePerson(app)
	return app
}
