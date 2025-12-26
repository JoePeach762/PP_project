package studentsService

import (
	"context"

	"github.com/JoePeach762/PP_project/internal/models"
)

func (s *StudentService) GetStudentInfoByIDs(ctx context.Context, IDs []uint64) ([]*models.StudentInfo, error) {
	return s.studentStorage.GetStudentInfoByIDs(ctx, IDs)
}
