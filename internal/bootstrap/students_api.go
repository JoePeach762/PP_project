package bootstrap

import (
	server "github.com/JoePeach762/PP_project/internal/api/student_service_api"
	"github.com/JoePeach762/PP_project/internal/services/studentsService"
)

func InitStudentServiceAPI(studentService *studentsService.StudentService) *server.StudentServiceAPI {
	return server.NewStudentServiceAPI(studentService)
}
