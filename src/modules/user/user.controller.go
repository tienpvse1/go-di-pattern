package user

import (
	"fmt"
	"tienpvse1/go-fiber-server/src/common"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service *UserService `inject:"UserService"`
	Router  *common.Router
}

func (controller UserController) Create() {
	b := controller.Router
	b.SetPrefix("/user")
	// alias service to later usage
	service := controller.Service

	b.Get("")
	b.Handler(func(c *fiber.Ctx) error {
		authors, err := service.FindOne()
		if err != nil {
			fmt.Print(err)
			return c.JSON(fiber.ErrBadRequest)
		}
		return c.JSON(authors)
	})

	b.Post("")
  b.AddMiddlewares()
	b.Handler(func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "this route post the user",
		})
	})
}
