package rest

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-fiber-demo/pkg/user"
)

func getAllUsers(u user.Service) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		users := u.GetAllUsers()
		c.JSON(users)
	}
}

func getUser(u user.Service) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		id := extractId(c)
		user := u.GetUser(id)
		c.JSON(user)
	}
}

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

func deleteUser(u user.Service) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		id := extractId(c)
		u.DeleteUser(id)
		c.JSON(fiber.Map{
			"ok": true,
		})
	}
}

func extractID(c *fiber.Ctx) int {
	param := c.Params("id")
	id, _ := strconv.Atoi(param) // todo: error handling
	return id
}
