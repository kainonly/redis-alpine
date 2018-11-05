package facade

import (
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

func Register() {
	env := config.NewConfig()
	env.Load(file.NewSource(
		file.WithPath("./env.yaml"),
	))
	options := env.Map()
	// redis facade
	if options["redis"] != nil {
		useRedis(options["redis"].(map[string]interface{}))
	}
}
