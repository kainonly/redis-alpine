package facade

import (
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

var configs map[string]interface{}

func Register() {
	c := config.NewConfig()
	c.Load(file.NewSource(
		file.WithPath("./config.yaml"),
	))
	configs = c.Map()
	// init redis facade
	if configs["redis"] != nil {
		useRedis(configs["redis"].(map[string]interface{}))
	}
}
