package main

import (
	"framework/cache"
	"framework/facade"
	"github.com/go-redis/redis"
)

func main() {
	facade.Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	val, err := cache.GetAdmin()
	if err != nil {
		println(err.Error())
	}

	println(val)
}
