package facade

import (
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

func Register() {
	source := config.NewConfig()
	source.Load(file.NewSource(
		file.WithPath("./config.yaml"),
	))
	configs := source.Map()
	// init redis facade
	if configs["redis"] != nil {
		useRedis(configs["redis"].(map[string]interface{}))
	}
}
