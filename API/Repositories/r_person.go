package Repositories

import "github.com/gofiber/fiber/v2"

type person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Addrees string `json:"address"`
	Job     string `json:"job"`
}

func GetPersonWithObjectStatic(c *fiber.Ctx) error {
	Person := person{
		Name:    "Muh Syamsul",
		Age:     20,
		Addrees: "Jalan Jambu Mente",
		Job:     "Fullstack Golang",
	}
	return c.JSON(Person)
}

func GetPersonWithObjectBody(c *fiber.Ctx) error {
	return c.SendString("get person with object params")
}

func GetPersonWithQueryParams(c *fiber.Ctx) error {
	return c.SendString("get person with query params")
}
