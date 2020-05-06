package route

import (
	"github.com/gofiber/fiber"
	"github.com/patipolst/go-fiber-demo/controller"
)

type UserRoute struct {
	controller *controller.UserController
}

func NewUserRoute(controller *controller.UserController) *UserRoute {
	return &UserRoute{controller}
}

func (r *UserRoute) SetupUserRoutes(app *fiber.App) {
	group := app.Group("/v1")
	r.HandleUser(group)
}

func (r *UserRoute) HandleUser(group *fiber.Group) {
	group.Get("/users", r.controller.GetUsers)
	group.Get("/users/:id", r.controller.GetUser)
	group.Post("/users", r.controller.CreateUser)
	group.Delete("/users/:id", r.controller.DeleteUser)
}
