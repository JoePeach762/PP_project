package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/JoePeach762/PP_project/user_service/config"
	"github.com/JoePeach762/PP_project/user_service/internal/bootstrap"
	"github.com/JoePeach762/PP_project/user_service/internal/services/user"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yaml"
	}
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Не удалось загрузить конфигурацию: %v", err)
	}
	userStorage, statsStorage, err := bootstrap.InitPGStorage(cfg)
	if err != nil {
		log.Fatalf("Не удалось инициализировать pgstorage: %v", err)
	}
	kafkaProducer := bootstrap.InitKafkaProducer(cfg)
	defer func() {
		if err := kafkaProducer.Close(); err != nil {
			log.Printf("Не удалось закрыть Kafka producer сервиса пользователей: %v", err)
		}
	}()
	userService := bootstrap.InitUserService(userStorage, statsStorage, cfg)
	userGRPC := user.NewGRPCServer(userService)
	userProcessor := bootstrap.InitUserProcessor(userService)
	userConsumer := bootstrap.InitUserConsumer(cfg, userProcessor)
	outboxPublisher := bootstrap.InitOutboxPublisher(userStorage, kafkaProducer)
	server := bootstrap.NewServer()
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go outboxPublisher.Run(ctx)

	if err := server.AppRun(ctx, userGRPC, userConsumer, cfg.HTTPPort, cfg.GRPCPort); err != nil {
		log.Fatalf("Сервер пользователей завершился с ошибкой: %v", err)
	}
}
