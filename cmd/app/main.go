package main

import (
	"log"

	"github.com/JoePeach762/PP_project/config"
	"github.com/JoePeach762/PP_project/internal/bootstrap"
)

func main() {

	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal("failed to load config:", err)
	}

	pgConn := cfg.Database.ConnString()
	redisAddr := cfg.Redis.Addr
	userAgent := cfg.MealServiceSettings.OFFUserAgent

	studentsStorage := bootstrap.InitPGStorage(cfg)
	studentService := bootstrap.InitStudentService(studentsStorage, cfg)
	studentsInfoProcessor := bootstrap.InitStudentsInfoProcessor(studentService)
	studentsinfoupsertconsumer := bootstrap.InitStudentInfoUpsertConsumer(cfg, studentsInfoProcessor)
	studentsApi := bootstrap.InitStudentServiceAPI(studentService)

	bootstrap.AppRun(*studentsApi, studentsinfoupsertconsumer)
}
