package main

import (
	"log"
	"os"

	"github.com/JoePeach762/PP_project/user_service/config"
	"github.com/JoePeach762/PP_project/user_service/internal/bootstrap"
	"github.com/JoePeach762/PP_project/user_service/internal/services/user"
)

func main() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "config.yaml"
	}

	cfg, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		config.LoadConfig(configPath)
	}

	pgStorage, err := bootstrap.InitPGStorage(cfg)
	if err != nil {
		log.Fatalf("Failed to load pgstorage: %v", err)
	}

	userService := bootstrap.InitUserService(pgStorage, cfg)

	userGRPC := user.NewGRPCServer(userService)

	userProcessor := bootstrap.InitUserProcessor(userService)
	userConsumer := bootstrap.InitUserConsumer(cfg, userProcessor)

	server := bootstrap.NewServer()
	if err := server.AppRun(userGRPC, userConsumer); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
