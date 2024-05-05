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
	routes.Post("/person/body", Controller.PersonWithObjectBody)
	routes.Post("/person/body/all", Controller.PersonWithObjectBodyAll)
	routes.Get("/person/query", Controller.PersonWithObjectQuery)
	routes.Get("/person/query/index", Controller.PersonWithObjectQueryByIndex)
	routes.Get("/person/params", Controller.PersonWithObjectParams)
}

func Setup() *fiber.App {
	fastergoding.Run()
	app := fiber.New()
	SetupRoutePerson(app)
	return app
}
