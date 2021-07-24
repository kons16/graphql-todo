package redis

import "github.com/go-redis/redis/v8"

func NewRedisDB() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return rdb, nil
}
