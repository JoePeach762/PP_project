package user

import (
	"context"
	"errors"
	"testing"

	"github.com/JoePeach762/PP_project/user_service/internal/services/user/mocks"
)

func TestServiceDeleteByIds_DelegatesToStorage(t *testing.T) {
	ctx := context.Background()
	userStorage := &mocks.UserStorage{}
	service := newServiceForTest(userStorage, nil)
	ids := []uint64{1, 2, 3}
	expectedErr := errors.New("ошибка удаления")

	userStorage.EXPECT().DeleteUsersAndEnqueueEvent(ctx, ids).Return(expectedErr)

	err := service.DeleteByIds(ctx, ids)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("ожидалась ошибка %v, получено %v", expectedErr, err)
	}

	userStorage.AssertExpectations(t)
}
