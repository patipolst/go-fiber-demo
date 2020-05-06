package main

import (
	"github.com/gofiber/fiber"
	"github.com/patipolst/go-fiber-demo/service"
)

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
}

func helloHandler(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloHandler)

	app.Get("/users", service.GetUsers)
	app.Get("/users/:id", service.GetUser)
	app.Post("/users", service.CreateUser)
	app.Delete("/users/:id", service.DeleteUser)
}
