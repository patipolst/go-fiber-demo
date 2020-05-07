package rest

import (
	"github.com/gofiber/fiber"
	"github.com/patipolst/go-fiber-demo/pkg/user"
)

func Handler(u user.Service) *fiber.App {
	app := fiber.New()

	group := app.Group("/v1")

	group.Get("/users", getAllUsers(u))

	return app
}

// addBeer returns a handler for POST /beers requests
func getAllUsers(u user.Service) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		u.GetAllUsers()
		c.Send("Hello, Add user!")
	}
}
