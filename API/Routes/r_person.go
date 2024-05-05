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

	//version app 1.0
	routes := app.Group("/api")
	routes.Get("/person/static", Controller.PersonWithObjectStatic)
	routes.Get("/person/body", Controller.PersonWithObjectBody)
	routes.Get("/person/query", Controller.PersonWithQueryParams)
}

func Setup() *fiber.App {
	fastergoding.Run()
	app := fiber.New()
	SetupRoutePerson(app)
	return app
}
