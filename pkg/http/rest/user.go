package rest

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-fiber-demo/pkg/user"
)

type UserHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) UserHandler {
	return UserHandler{service}
}

func (handler *UserHandler) getAllUsers(c *fiber.Ctx) {
	users := handler.service.GetAllUsers()
	c.JSON(users)
}

func (handler *UserHandler) getUser(c *fiber.Ctx) {
	id := extractID(c)
	user := handler.service.GetUser(id)
	c.JSON(user)
}

func (handler *UserHandler) createUser(c *fiber.Ctx) {
	user := new(user.User)
	if err := c.BodyParser(user); err != nil {
		log.Fatal(err)
	}
	handler.service.CreateUser(*user)
	c.JSON(user)
}

func (handler *UserHandler) deleteUser(c *fiber.Ctx) {
	id := extractID(c)
	handler.service.DeleteUser(id)
	c.JSON(fiber.Map{
		"ok": true,
	})
}

func extractID(c *fiber.Ctx) int {
	param := c.Params("id")
	id, _ := strconv.Atoi(param) // todo: error handling
	return id
}
