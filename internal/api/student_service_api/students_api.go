package student_service_api

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
	"github.com/JoePeach762/PP_project/internal/pb/students_api"
)

type studentService interface {
	GetStudentInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.StudentInfo, error)
	UpsertStudentInfo(ctx context.Context, studentsInfos []*models.StudentInfo) error
}

// StudentServiceAPI реализует grpc StudentsServiceServer
type StudentServiceAPI struct {
	students_api.UnimplementedStudentsServiceServer
	studentService studentService
}

func NewStudentServiceAPI(studentService studentService) *StudentServiceAPI {
	return &StudentServiceAPI{
		studentService: studentService,
	}
}
