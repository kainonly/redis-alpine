package main

import (
	"framework/api"
	"framework/facade"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	facade.Register()
	api.Register(app)

	app.Run(iris.Addr(":80"))
}
