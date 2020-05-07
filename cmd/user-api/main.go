package main

import (
	"github.com/patipolst/go-fiber-demo/pkg/http/rest"
	"github.com/patipolst/go-fiber-demo/pkg/store/stub"
	"github.com/patipolst/go-fiber-demo/pkg/user"
)

func main() {
	userStore := stub.NewStore()
	userService := user.NewService(userStore)
	app := rest.Handler(userService)
	app.Listen(3000)
}

// func helloHandler(c *fiber.Ctx) {
// 	c.Send("Hello, World!")
// }

// func setupRoutes(app *fiber.App) {
// 	app.Get("/", helloHandler)

// 	userStore := user.NewDummyStore()
// 	userService := user.NewService(userStore)
// 	userController := user.NewController(userService)
// 	userRoute := user.NewRoute(userController)
// 	userRoute.SetupRoutes(app)
// }
