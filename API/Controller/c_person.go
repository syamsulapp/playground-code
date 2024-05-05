package Controller

import (
	"idiJakartaPusat/Repositories"

	"github.com/gofiber/fiber/v2"
)

func PersonWithObjectStatic(c *fiber.Ctx) error {
	return Repositories.GetPersonWithObjectStatic(c)
}

func PersonWithBodyParams(c *fiber.Ctx) error {
	return Repositories.GetPersonWithBodyParams(c)
}

func PersonWithQueryParams(c *fiber.Ctx) error {
	return Repositories.GetPersonWithQueryParams(c)
}

func PersonWithQueryParamsByIndex(c *fiber.Ctx) error {
	return Repositories.GetPersonWithQueryParamsByIndex(c)
}

func PersonWithBodyParamsAll(c *fiber.Ctx) error {
	return Repositories.GetPersonWithBodyParamsAll(c)
}

func PersonWithRouteParams(c *fiber.Ctx) error {
	return Repositories.GetPersonWithRouteParams(c)
}
