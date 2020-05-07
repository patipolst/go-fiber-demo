package rest

import (
	"github.com/gofiber/fiber"
	"github.com/patipolst/go-fiber-demo/pkg/user"
)

// Handler returns a fiber instance with defined handlers
func Handler(u user.Service) *fiber.App {
	app := fiber.New()

	v1 := app.Group("/v1")
	v1.Get("/users", getAllUsers(u))
	v1.Get("/users/:id", getUser(u))
	v1.Post("/users", createUser(u))
	v1.Delete("/users/:id", deleteUser(u))

	return app
}
