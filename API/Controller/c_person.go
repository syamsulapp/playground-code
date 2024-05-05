package Controller

import (
	"idiJakartaPusat/Repositories"

	"github.com/gofiber/fiber/v2"
)

func PersonWithObjectStatic(c *fiber.Ctx) error {
	return Repositories.GetPersonWithObjectStatic(c)
}

func PersonWithObjectBody(c *fiber.Ctx) error {
	return Repositories.GetPersonWithObjectBody(c)
}

func PersonWithObjectQuery(c *fiber.Ctx) error {
	return Repositories.GetPersonWithObjectQuery(c)
}

func PersonWithObjectQueryByIndex(c *fiber.Ctx) error {
	return Repositories.GetPersonWithObjectQueryByIndex(c)
}

func PersonWithObjectBodyAll(c *fiber.Ctx) error {
	return Repositories.GetPersonWithObjectBodyAll(c)
}

func PersonWithObjectParams(c *fiber.Ctx) error {
	return Repositories.GetPersonWithObjectParams(c)
}
