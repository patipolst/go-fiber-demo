package controller

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/patipolst/go-fiber-demo/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{service}
}

func (c *UserController) GetUsers(ctx *fiber.Ctx) {
	users := c.service.GetUsers()
	ctx.JSON(users)
}

func (c *UserController) GetUser(ctx *fiber.Ctx) {
	param := ctx.Params("id")
	id, _ := strconv.Atoi(param) // todo: error handling
	user := c.service.GetUser(id)
	ctx.JSON(user)
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) {
	user := new(service.User)
	if err := ctx.BodyParser(user); err != nil {
		log.Fatal(err)
	}
	users := c.service.CreateUser(*user)
	ctx.JSON(users)
}

func (c *UserController) DeleteUser(ctx *fiber.Ctx) {
	param := ctx.Params("id")
	id, _ := strconv.Atoi(param) // todo: error handling
	c.service.DeleteUser(id)
	ctx.JSON(fiber.Map{
		"ok": true,
	})
}
