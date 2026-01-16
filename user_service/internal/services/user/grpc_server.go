package user

import (
	"context"

	models "github.com/JoePeach762/PP_project/user_service/internal/models"
	pbmodels "github.com/JoePeach762/PP_project/user_service/internal/pb/models/user_models"
	"github.com/JoePeach762/PP_project/user_service/internal/pb/users_api"
)

type GRPCServer struct {
	users_api.UnimplementedUserServiceServer
	service *Service
}

func NewGRPCServer(service *Service) *GRPCServer {
	return &GRPCServer{service: service}
}

func (s *GRPCServer) CreateUsers(ctx context.Context, req *users_api.CreateUsersRequest) (*users_api.CreateUsersResponse, error) {
	var users []*models.UserInfo
	for _, u := range req.Users {
		users = append(users, &models.UserInfo{
			ID:             u.Id,
			Name:           u.Name,
			Email:          u.Email,
			Sex:            u.Sex,
			Age:            uint32(u.Age),
			HeightCm:       uint32(u.HeightCm),
			WeightKg:       uint32(u.WeightKg),
			TargetWeightKg: uint32(u.TargetWeightKg),
		})
	}

	if err := s.service.Add(ctx, users); err != nil {
		return nil, err
	}

	return &users_api.CreateUsersResponse{}, nil
}

func (s *GRPCServer) UpdateUser(ctx context.Context, req *users_api.UpdateUserRequest) (*users_api.UpdateUserResponse, error) {
	u := req.User
	user := models.UserInfo{
		ID:             u.Id,
		Name:           u.Name,
		Email:          u.Email,
		Sex:            u.Sex,
		Age:            uint32(u.Age),
		HeightCm:       uint32(u.HeightCm),
		WeightKg:       uint32(u.WeightKg),
		TargetWeightKg: uint32(u.TargetWeightKg),
	}

	if err := s.service.Update(ctx, u.Id, user); err != nil {
		return nil, err
	}

	return &users_api.UpdateUserResponse{}, nil
}

func (s *GRPCServer) GetUsers(ctx context.Context, req *users_api.GetUsersRequest) (*users_api.GetUsersResponse, error) {
	users, err := s.service.GetByIds(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	var pbUsers []*pbmodels.UserInfo
	for _, u := range users {
		pbUsers = append(pbUsers, &pbmodels.UserInfo{
			Id:              u.ID,
			Name:            u.Name,
			Email:           u.Email,
			Sex:             u.Sex,
			Age:             uint32(u.Age),
			HeightCm:        uint32(u.HeightCm),
			WeightKg:        uint32(u.WeightKg),
			TargetWeightKg:  uint32(u.TargetWeightKg),
			CurrentCalories: uint32(u.CurrentCalories),
			CurrentProteins: uint32(u.CurrentProteins),
			CurrentFats:     uint32(u.CurrentFats),
			CurrentCarbs:    uint32(u.CurrentCarbs),
			TargetCalories:  uint32(u.TargetCalories),
			TargetProteins:  uint32(u.TargetProteins),
			TargetFats:      uint32(u.TargetFats),
			TargetCarbs:     uint32(u.TargetCarbs),
		})
	}

	return &users_api.GetUsersResponse{Users: pbUsers}, nil
}

func (s *GRPCServer) DeleteUsers(ctx context.Context, req *users_api.DeleteUsersRequest) (*users_api.DeleteUsersResponse, error) {
	if err := s.service.DeleteByIds(ctx, req.Ids); err != nil {
		return nil, err
	}
	return &users_api.DeleteUsersResponse{}, nil
}
