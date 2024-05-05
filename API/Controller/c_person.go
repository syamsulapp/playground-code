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

func PersonWithQueryParams(c *fiber.Ctx) error {
	return Repositories.GetPersonWithQueryParams(c)
}
