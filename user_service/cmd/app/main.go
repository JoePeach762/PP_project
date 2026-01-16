package main

import (
	"log"
	"os"

	"github.com/JoePeach762/PP_project/user_service/config"
	"github.com/JoePeach762/PP_project/user_service/internal/bootstrap"
	"github.com/JoePeach762/PP_project/user_service/internal/services/user"
)

func main() {
	cfg, err := config.LoadConfig(os.Getenv("CONFIG_PATH_USER"))
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Инициализация зависимостей
	pgStorage, err := bootstrap.InitPGStorage(cfg)
	if err != nil {
		log.Fatalf("Failed to load pgstorage: %v", err)
	}

	// Сервисы
	userService := bootstrap.InitUserService(pgStorage, cfg)

	// gRPC-серверы
	userGRPC := user.NewGRPCServer(userService)

	// Kafka
	userProcessor := bootstrap.InitUserProcessor(userService)
	userConsumer := bootstrap.InitUserConsumer(cfg, userProcessor)

	// Запуск сервера
	server := bootstrap.NewServer()
	if err := server.AppRun(userGRPC, userConsumer); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
