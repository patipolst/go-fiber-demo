package user

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{service}
}

func (c *Controller) GetAll(ctx *fiber.Ctx) {
	models := c.service.GetAll()
	ctx.JSON(models)
}

func (c *Controller) Get(ctx *fiber.Ctx) {
	param := ctx.Params("id")
	id, _ := strconv.Atoi(param) // todo: error handling
	model := c.service.Get(id)
	ctx.JSON(model)
}

func (c *Controller) Create(ctx *fiber.Ctx) {
	model := new(Model)
	if err := ctx.BodyParser(model); err != nil {
		log.Fatal(err)
	}
	models := c.service.Create(*model)
	ctx.JSON(models)
}

func (c *Controller) Delete(ctx *fiber.Ctx) {
	param := ctx.Params("id")
	id, _ := strconv.Atoi(param) // todo: error handling
	c.service.Delete(id)
	ctx.JSON(fiber.Map{
		"ok": true,
	})
}
