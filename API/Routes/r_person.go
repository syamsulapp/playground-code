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
	routes.Post("/person/body", Controller.PersonWithBodyParams)
	routes.Post("/person/body/all", Controller.PersonWithBodyParamsAll)
	routes.Get("/person/query", Controller.PersonWithQueryParams)
	routes.Get("/person/query/index", Controller.PersonWithQueryParamsByIndex)
	routes.Get("/person/params/:name/:email/:username", Controller.PersonWithRouteParams)
}

func Setup() *fiber.App {
	fastergoding.Run()
	app := fiber.New()
	SetupRoutePerson(app)
	return app
}
