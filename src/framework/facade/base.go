package facade

import (
	"github.com/micro/go-config"
	"github.com/micro/go-config/source/file"
)

var configs map[string]interface{}

func Config(key string) string {
	return configs[key].(string)
}

func Register() {
	c := config.NewConfig()
	c.Load(file.NewSource(
		file.WithPath("./config.json"),
	))
	configs = c.Map()

	if configs["redis"] != nil {
		useRedis(configs["redis"])
	}
}
