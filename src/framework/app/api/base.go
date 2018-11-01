package api

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func Set(app *iris.Application) {
	m := mvc.New(app.Party("api",
		cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowCredentials: true,
		}),
	))
	m.Handle(new(IndexController))
}
