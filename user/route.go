// package user

// import (
// 	"github.com/gofiber/fiber"
// )

// type Route struct {
// 	controller *Controller
// }

// func NewRoute(controller *Controller) *Route {
// 	return &Route{controller}
// }

// func (r *Route) SetupRoutes(app *fiber.App) {
// 	group := app.Group("/v1")
// 	r.Handle(group)
// }

// func (r *Route) Handle(group *fiber.Group) {
// 	group.Get("/users", r.controller.GetAll)
// 	group.Get("/users/:id", r.controller.Get)
// 	group.Post("/users", r.controller.Create)
// 	group.Delete("/users/:id", r.controller.Delete)
// }
