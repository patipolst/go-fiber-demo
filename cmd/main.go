package main

import "github.com/gofiber/fiber"

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
}
