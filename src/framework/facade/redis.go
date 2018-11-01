package facade

import (
	"github.com/go-redis/redis"
)

var Redis *redis.Client

func useRedis(options interface{}) {
	ops := options.(map[string]interface{})
	Redis = redis.NewClient(&redis.Options{
		Addr:     ops["address"].(string),
		Password: ops["password"].(string),
		DB:       int(ops["db"].(byte)),
	})
}
