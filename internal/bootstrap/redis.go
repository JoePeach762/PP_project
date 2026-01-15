package bootstrap

import (
	"github.com/JoePeach762/PP_project/config"
	redisstore "github.com/JoePeach762/PP_project/internal/storage/redis"
	"github.com/redis/go-redis/v9"
)

func InitRedisCache(cfg *config.Config) (*redisstore.RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	return redisstore.NewRedisCache(client)
}
