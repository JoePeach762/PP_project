package main

import (
	"fmt"
	"os"

	"github.com/JoePeach762/PP_project/config"
	"github.com/JoePeach762/PP_project/internal/bootstrap"
)

func main() {

	configPath := os.Getenv("configPath")
	if configPath == "" {
		configPath = "config.yaml"
	}

	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		panic(fmt.Sprintf("ошибка парсинга конфига, %v", err))
	}

	studentsStorage := bootstrap.InitPGStorage(cfg)
	studentService := bootstrap.InitStudentService(studentsStorage, cfg)
	studentsInfoProcessor := bootstrap.InitStudentsInfoProcessor(studentService)
	studentsinfoupsertconsumer := bootstrap.InitStudentInfoUpsertConsumer(cfg, studentsInfoProcessor)
	studentsApi := bootstrap.InitStudentServiceAPI(studentService)

	bootstrap.AppRun(*studentsApi, studentsinfoupsertconsumer)
}
