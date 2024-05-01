package controller

import (
	"github.com/gofiber/fiber/v2"
)

type Person struct {
	FullName string
	Age      int
	Address  string
	Job      string
}

func GetPersonWithObjectStatic(c *fiber.Ctx) error {
	user := Person{
		FullName: "Muh samsul",
		Age:      20,
		Address:  "jalan jambu mente",
		Job:      "fullstack developer",
	}
	return c.JSON(user)
}

func GetPersonWithObjectBody(c *fiber.Ctx) error {
	return c.SendString("halo body object")
}

func GetPersonWithQueryParams(c *fiber.Ctx) error {
	return c.SendString("halo params object")
}
