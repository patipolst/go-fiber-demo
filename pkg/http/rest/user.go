package rest

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-fiber-demo/pkg/user"
)

type UserHandler struct {
	s user.Service
}

func NewUserHandler(s user.Service) UserHandler {
	return UserHandler{s}
}

func (h *UserHandler) getAllUsers(c *fiber.Ctx) {
	users := h.s.GetAllUsers()
	c.JSON(users)
}

func (h *UserHandler) getUser(c *fiber.Ctx) {
	id := extractID(c)
	user := h.s.GetUser(id)
	c.JSON(user)
}

func (h *UserHandler) createUser(c *fiber.Ctx) {
	user := new(user.User)
	if err := c.BodyParser(user); err != nil {
		log.Fatal(err)
	}
	h.s.CreateUser(*user)
	c.JSON(user)
}

func (h *UserHandler) deleteUser(c *fiber.Ctx) {
	id := extractID(c)
	h.s.DeleteUser(id)
	c.JSON(fiber.Map{
		"ok": true,
	})
}

func extractID(c *fiber.Ctx) int {
	param := c.Params("id")
	id, _ := strconv.Atoi(param) // todo: error handling
	return id
}
