package facade

import (
	"github.com/go-redis/redis"
	"strconv"
)

var Redis *redis.Client

func useRedis(ops map[string]interface{}) {
	db, err := strconv.Atoi(ops["db"].(string))
	if err != nil {
		return
	}
	Redis = redis.NewClient(&redis.Options{
		Addr:     ops["host"].(string),
		Password: ops["password"].(string),
		DB:       db,
	})
}
