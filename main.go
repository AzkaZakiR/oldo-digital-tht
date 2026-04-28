package main

import (
	"github.com/gofiber/fiber/v3"
)
type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
func main(){
	app := fiber.New()

	app.Get("/", func (c fiber.Ctx) error  {
		res := Response{
			Message: "Hello from Gofibber",
			Status:  200,
		}
		return c.JSON(res)
	})

	app.Listen(":3000")
}