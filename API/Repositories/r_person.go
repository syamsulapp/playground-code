package Repositories

import (
	"github.com/gofiber/fiber/v2"
)

type person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Addrees string `json:"address"`
	Job     string `json:"job"`
}

type PersonWithBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PersonWithQuery struct {
	IPK      string `json:"ipk"`
	Fakultas string `json:"fakultas"`
	Jurusan  string `json:"jurusan"`
	Prestasi string `json:"prestasi"`
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

func GetPersonWithBodyParams(c *fiber.Ctx) error {
	personWithBody := new(PersonWithBody)

	if err := c.BodyParser(personWithBody); err != nil {
		return err
	}

	return c.JSON(personWithBody)
}

func GetPersonWithBodyParamsAll(c *fiber.Ctx) error {
	return c.Send(c.Body())
}

func GetPersonWithQueryParams(c *fiber.Ctx) error {
	personWithQuery := new(PersonWithQuery)

	if err := c.QueryParser(personWithQuery); err != nil {
		return err
	}

	return c.JSON(personWithQuery)
}

func GetPersonWithQueryParamsByIndex(c *fiber.Ctx) error {
	personWithQueryIndex := PersonWithQuery{
		IPK:      c.Query("ipk"),
		Fakultas: c.Query("fakultas"),
		Jurusan:  c.Query("jurusan"),
		Prestasi: c.Query("prestasi"),
	}
	return c.JSON(personWithQueryIndex)
}

func GetPersonWithRouteParams(c *fiber.Ctx) error {
	personWithRouteParams := PersonWithBody{
		Name:     c.Params("name"),
		Email:    c.Params("email"),
		Username: c.Params("username"),
	}
	return c.JSON(personWithRouteParams)
}
