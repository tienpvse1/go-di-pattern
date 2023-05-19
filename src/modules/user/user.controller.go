package user

import (
	"tienpvse1/go-fiber-server/src/common"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
  Service *UserService `inject:"UserService"`
  Router *common.Router
}


func (controller UserController) Create(){
  b := controller.Router
  service := controller.Service
  
  b.SetPrefix("/user")

  b.Get("")
  b.Handler(func(c *fiber.Ctx) error {
    authors, err := service.FindOne()
    if err != nil {
      return c.JSON(fiber.ErrBadRequest)
    }
    return c.JSON(authors)
  })
}

