package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetPerson(c *fiber.Ctx) error {
	return c.SendString("halo get person")
}
