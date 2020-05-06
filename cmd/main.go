package main

import (
	"github.com/gofiber/fiber"
	"github.com/patipolst/go-fiber-demo/controller"
	"github.com/patipolst/go-fiber-demo/route"
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

	userStore := service.NewDummyUserStore()
	userService := service.NewUserService(userStore)
	userController := controller.NewUserController(userService)
	userRoute := route.NewUserRoute(userController)
	userRoute.SetupUserRoutes(app)
}
