package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

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
	defer func() {
		if err := kafkaProducer.Close(); err != nil {
			log.Printf("Failed to close meal Kafka producer: %v", err)
		}
	}()
	mealService := bootstrap.InitMealService(pgStorage, redisCache, offClient, cfg)
	mealGRPC := meal.NewGRPCServer(mealService)
	mealProcessor := bootstrap.InitMealProcessor(mealService)
	mealConsumer := bootstrap.InitMealConsumer(cfg, mealProcessor)
	outboxPublisher := bootstrap.InitOutboxPublisher(pgStorage, kafkaProducer)
	server := bootstrap.NewServer()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go outboxPublisher.Run(ctx)

	if err := server.AppRun(ctx, mealGRPC, mealConsumer, cfg.HTTPPort, cfg.GRPCPort); err != nil {
		log.Fatalf("meal server failed: %v", err)
	}
}
