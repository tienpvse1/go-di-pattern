package main

import (
	"log"
	database "tienpvse1/go-fiber-server/src"
	"tienpvse1/go-fiber-server/src/common"
	"tienpvse1/go-fiber-server/src/modules/user"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type User struct {
	name           string
	SomeOtherField bool
}

func main() {
	app := fiber.New()
	queries, err := database.InitDatabase()
	if err != nil {
		panic("cannot init database connection ")
	}

	module := common.Bundler{
		Queries: queries,
    Router: common.RouteBuilder{App: app}.CreateRouteBuilder(),
		Controllers: []common.IController{
			new(user.UserController),
		},
		Services: []interface{}{
			new(user.UserService),
		},
	}
	module.Bundle()
	log.Fatal(app.Listen(":3000"))

}
