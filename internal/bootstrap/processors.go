package bootstrap

import (
	studentsinfoprocessor "github.com/JoePeach762/PP_project/internal/services/processors/students_info_processor"
	"github.com/JoePeach762/PP_project/internal/services/studentsService"
)

func InitStudentsInfoProcessor(studentService *studentsService.StudentService) *studentsinfoprocessor.StudentsInfoProcessor {
	return studentsinfoprocessor.NewStudentsInfoProcessor(studentService)
}
