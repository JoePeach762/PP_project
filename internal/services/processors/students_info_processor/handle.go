package studentsinfoprocessor

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (p *StudentsInfoProcessor) Handle(ctx context.Context, studentsInfo *models.StudentInfo) error {
	return p.studentService.UpsertStudentInfo(ctx, []*models.StudentInfo{studentsInfo})
}
