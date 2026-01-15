package bootstrap

import (
	"context"

	userprocessor "github.com/JoePeach762/PP_project/internal/services/processors/user"
	"github.com/JoePeach762/PP_project/internal/services/user"
)

func InitUserProcessor(userService *user.Service) *userprocessor.Processor {
	return userprocessor.NewUserProcessor(context.Background(), userService)
}
