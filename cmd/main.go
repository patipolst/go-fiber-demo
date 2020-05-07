package main

import (
	"github.com/gofiber/fiber"
	"github.com/patipolst/go-fiber-demo/user"
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

	userStore := user.NewDummyStore()
	userService := user.NewService(userStore)
	userController := user.NewController(userService)
	userRoute := user.NewRoute(userController)
	userRoute.SetupRoutes(app)
}
