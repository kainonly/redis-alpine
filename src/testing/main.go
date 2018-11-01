package main

import (
	"framework/facade"
)

func main() {
	facade.Register()
	println(facade.Config("name"))
}
