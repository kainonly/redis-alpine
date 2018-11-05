package main

import (
	"framework/cache"
	"framework/facade"
)

func main() {
	facade.Register()
	str, err := cache.GetAdmin()
	if err != nil {
		println(err.Error())
	}
	println(str)
}
