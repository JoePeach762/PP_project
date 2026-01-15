package meal

import (
	"context"

	models "github.com/JoePeach762/PP_project/internal/models"
	"github.com/JoePeach762/PP_project/internal/pb/meals_api"
	pbmodels "github.com/JoePeach762/PP_project/internal/pb/models"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCServer struct {
	meals_api.UnimplementedMealServiceServer
	service *Service
}

func NewGRPCServer(service *Service) *GRPCServer {
	return &GRPCServer{service: service}
}

func (s *GRPCServer) AddMeal(ctx context.Context, req *meals_api.AddMealRequest) (*emptypb.Empty, error) {
	mealInput := &models.MealInput{
		UserID:      req.UserId,
		Name:        req.Name,
		WeightGrams: req.WeightGrams,
	}

	if err := s.service.Add(ctx, mealInput); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (s *GRPCServer) GetMeals(ctx context.Context, req *meals_api.GetMealsRequest) (*meals_api.GetMealsResponse, error) {
	meals, err := s.service.GetMealsByUserId(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	var pbMeals []*pbmodels.MealInfo
	for _, m := range meals {
		pbMeals = append(pbMeals, &pbmodels.MealInfo{
			Id:            m.ID,
			UserId:        m.UserId,
			Name:          m.Name,
			WeightGrams:   m.WeightGrams,
			Calories_100G: m.Calories100g,
			Proteins_100G: m.Proteins100g,
			Fats_100G:     m.Fats100g,
			Carbs_100G:    m.Carbs100g,
			Date:          timestamppb.New(m.Date),
		})
	}

	return &meals_api.GetMealsResponse{Meals: pbMeals}, nil
}
