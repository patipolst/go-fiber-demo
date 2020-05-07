package rest

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-fiber-demo/pkg/user"
)

func getAllUsers(s user.Service) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		users := s.GetAllUsers()
		c.JSON(users)
	}
}

func getUser(s user.Service) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		id := extractID(c)
		user := s.GetUser(id)
		c.JSON(user)
	}
}

func createUser(s user.Service) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		user := new(user.User)
		if err := c.BodyParser(user); err != nil {
			log.Fatal(err)
		}
		s.CreateUser(*user)
		c.JSON(user)
	}
}

func deleteUser(s user.Service) func(c *fiber.Ctx) {
	return func(c *fiber.Ctx) {
		id := extractID(c)
		s.DeleteUser(id)
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
