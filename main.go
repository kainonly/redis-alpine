package main

import (
	"framework/app/api"
	"framework/app/system"
	"framework/facade"
	"github.com/go-redis/redis"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())

	facade.Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	system.Set(app)
	api.Set(app)
	app.Run(iris.Addr(":80"))
}
