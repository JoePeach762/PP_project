package studentsService

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *StudentService) UpsertStudentInfo(ctx context.Context, studentsInfos []*models.StudentInfo) error {

	if err := s.validateInfo(studentsInfos); err != nil {
		return err
	}
	return s.studentStorage.UpsertStudentInfo(ctx, studentsInfos)
}
