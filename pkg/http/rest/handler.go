package rest

import (
	"github.com/gofiber/fiber"
)

// New returns a fiber instance with defined api handlers
func New(h UserHandler) *fiber.App {
	app := fiber.New()

	v1 := app.Group("/v1")
	v1.Get("/users", h.getAllUsers)
	v1.Get("/users/:id", h.getUser)
	v1.Post("/users", h.createUser)
	v1.Delete("/users/:id", h.deleteUser)

	return app
}
