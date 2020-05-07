package rest

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-fiber-demo/pkg/user"
)

func Handler(u user.Service) *fiber.App {
	app := fiber.New()

	group := app.Group("/v1")

	group.Get("/users", getAllUsers(u))
	group.Get("/users/:id", getUser(u))
	group.Post("/users", createUser(u))
	group.Delete("/users/:id", deleteUser(u))

	return app
}

// addBeer returns a handler for POST /beers requests
func getAllUsers(u user.Service) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		users := u.GetAllUsers()
		c.JSON(users)
	}
}

// addBeer returns a handler for POST /beers requests
func getUser(u user.Service) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		param := c.Params("id")
		id, _ := strconv.Atoi(param) // todo: error handling
		user := u.GetUser(id)
		c.JSON(user)
	}
}

// addBeer returns a handler for POST /beers requests
func createUser(u user.Service) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		user := new(user.User)
		if err := c.BodyParser(user); err != nil {
			log.Fatal(err)
		}
		u.CreateUser(*user)
		c.JSON(user)
	}
}

// addBeer returns a handler for POST /beers requests
func deleteUser(u user.Service) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		param := c.Params("id")
		id, _ := strconv.Atoi(param) // todo: error handling
		u.DeleteUser(id)
		c.JSON(fiber.Map{
			"ok": true,
		})
	}
}
