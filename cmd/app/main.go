package main

import (
	"log"
	"os"

	"github.com/JoePeach762/PP_project/config"
	"github.com/JoePeach762/PP_project/internal/bootstrap"
	"github.com/JoePeach762/PP_project/internal/services/meal"
	"github.com/JoePeach762/PP_project/internal/services/user"
)

func main() {
	cfg, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Инициализация зависимостей
	pgStorage, err := bootstrap.InitPGStorage(cfg)
	if err != nil {
		log.Fatalf("Failed to load pgstorage: %v", err)
	}
	redisCache, err := bootstrap.InitRedisCache(cfg)
	if err != nil {
		log.Fatalf("Failed to load rediscache: %v", err)
	}
	kafkaProducer := bootstrap.InitKafkaProducer(cfg)
	offClient := bootstrap.InitOFFClient(cfg)

	// Сервисы
	userService := bootstrap.InitUserService(pgStorage, cfg)
	mealService := bootstrap.InitMealService(pgStorage, redisCache, kafkaProducer, offClient, cfg)

	// gRPC-серверы
	userGRPC := user.NewGRPCServer(userService)
	mealGRPC := meal.NewGRPCServer(mealService)

	// Kafka
	userProcessor := bootstrap.InitUserProcessor(userService)
	userConsumer := bootstrap.InitUserConsumer(cfg, userProcessor)

	// Запуск сервера
	server := bootstrap.NewServer()
	if err := server.AppRun(userGRPC, mealGRPC, userConsumer); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
