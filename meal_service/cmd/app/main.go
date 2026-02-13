package main

import (
	"log"
	"os"

	"github.com/JoePeach762/PP_project/meal_service/config"
	"github.com/JoePeach762/PP_project/meal_service/internal/bootstrap"
	"github.com/JoePeach762/PP_project/meal_service/internal/services/meal"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yaml"
	}
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	pgStorage, err := bootstrap.InitPGStorage(cfg)
	if err != nil {
		log.Fatalf("Failed to load pgstorage: %v", err)
	}
	redisCache, err := bootstrap.InitRedisCache(cfg)
	if err != nil {
		log.Fatalf("Failed to load rediscache: %v", err)
	}
	offClient := bootstrap.InitOFFClient(cfg)
	kafkaProducer := bootstrap.InitKafkaProducer(cfg)
	mealService := bootstrap.InitMealService(pgStorage, redisCache, kafkaProducer, offClient, cfg)
	mealGRPC := meal.NewGRPCServer(mealService)
	mealProcessor := bootstrap.InitMealProcessor(mealService)
	mealConsumer := bootstrap.InitMealConsumer(cfg, mealProcessor)
	server := bootstrap.NewServer()
	if err := server.AppRun(mealGRPC, mealConsumer); err != nil {
		log.Fatalf("meal server failed: %v", err)
	}
}
