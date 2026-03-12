package meal

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log/slog"
	"time"

	"github.com/JoePeach762/PP_project/meal_service/internal/models"
)

func (s *Service) Add(ctx context.Context, req *models.MealInput) error {
	if err := s.validate(req); err != nil {
		return err
	}

	template, err := s.cache.GetProduct(ctx, req.Name)
	if err != nil {
		return fmt.Errorf("get product from cache: %w", err)
	}
	if template == nil {
		template, err = s.offClient.FetchProduct(ctx, req.Name)
		if err != nil {
			return err
		}
		if err := s.cache.AddProduct(ctx, template); err != nil {
			slog.Warn("Failed to cache product", "name", template.Name, "error", err)
		}
	}

	eventID, err := newEventID()
	if err != nil {
		return fmt.Errorf("generate meal event id: %w", err)
	}

	info := &models.MealInfo{
		EventID:      eventID,
		Name:         req.Name,
		UserId:       req.UserID,
		WeightGrams:  req.WeightGrams,
		Calories100g: template.Calories100g,
		Proteins100g: template.Proteins100g,
		Fats100g:     template.Fats100g,
		Carbs100g:    template.Carbs100g,
		Date:         time.Now(),
	}

	return s.storage.AddMealAndEnqueueEvent(ctx, info)
}

func newEventID() (string, error) {
	var raw [16]byte
	if _, err := rand.Read(raw[:]); err != nil {
		return "", err
	}

	return hex.EncodeToString(raw[:]), nil
}
