package bootstrap

import (
	"context"

	"github.com/JoePeach762/PP_project/config"
	"github.com/JoePeach762/PP_project/internal/services/studentsService"
	"github.com/JoePeach762/PP_project/internal/storage/pgstorage"
)

func InitStudentService(storage *pgstorage.PGstorage, cfg *config.Config) *studentsService.StudentService {

	return studentsService.NewStudentService(context.Background(), storage, cfg.StudentServiceSettings.MinNameLen, cfg.StudentServiceSettings.MaxNameLen)
}
