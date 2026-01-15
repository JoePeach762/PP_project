package bootstrap

import (
	RedisCache "github.com/JoePeach762/PP_project/internal/storage/redis"
	"github.com/redis/go-redis/v9"
)

func NewRedisCache(addr, password string, db int) (*RedisCache.RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return RedisCache.NewRedisCache(client)
}
