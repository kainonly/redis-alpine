package cache

import "framework/facade"

var key = "admin"

func FactoryAdmin() error {
	return facade.Redis.Set(key, "test", 0).Err()
}

func GetAdmin() (string, error) {
	return facade.Redis.Get(key).Result()
}
