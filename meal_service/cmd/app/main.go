package main

import (
	"log"
	"os"

	"github.com/JoePeach762/PP_project/meal_service/config"
	"github.com/JoePeach762/PP_project/meal_service/internal/bootstrap"
	"github.com/JoePeach762/PP_project/meal_service/internal/services/meal"
)

func main() {
	cfg, err := config.LoadConfig(os.Getenv("CONFIG_PATH_MEALS"))
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
	mealService := bootstrap.InitMealService(pgStorage, redisCache, kafkaProducer, offClient, cfg)

	// gRPC-серверы
	mealGRPC := meal.NewGRPCServer(mealService)

	// Запуск сервера
	server := bootstrap.NewServer()
	if err := server.AppRun(mealGRPC); err != nil {
		log.Fatalf("meal server failed: %v", err)
	}
}
