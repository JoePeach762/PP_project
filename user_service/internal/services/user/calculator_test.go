package user

import (
	"strings"
	"testing"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
)

func TestServiceCalculateTargetsSingle_SetsTargets(t *testing.T) {
	service := newServiceForTest(nil, nil)
	info := &models.UserInfo{
		Sex:            "female",
		Age:            28,
		HeightCm:       165,
		WeightKg:       60,
		TargetWeightKg: 65,
	}

	if err := service.calculateTargetsSingle(info); err != nil {
		t.Fatalf("calculateTargetsSingle returned error: %v", err)
	}

	if info.TargetCalories != 2129 {
		t.Fatalf("unexpected TargetCalories: %d", info.TargetCalories)
	}

	if info.TargetProteins != 130 {
		t.Fatalf("unexpected TargetProteins: %d", info.TargetProteins)
	}

	if info.TargetFats != 65 {
		t.Fatalf("unexpected TargetFats: %d", info.TargetFats)
	}

	if info.TargetCarbs != 256 {
		t.Fatalf("unexpected TargetCarbs: %d", info.TargetCarbs)
	}
}

func TestServiceCalculateTargetsSingle_AppliesMinimumCalories(t *testing.T) {
	service := newServiceForTest(nil, nil)
	info := &models.UserInfo{
		Sex:            "female",
		Age:            18,
		HeightCm:       140,
		WeightKg:       40,
		TargetWeightKg: 35,
	}

	if err := service.calculateTargetsSingle(info); err != nil {
		t.Fatalf("calculateTargetsSingle returned error: %v", err)
	}

	if info.TargetCalories != 1200 {
		t.Fatalf("expected minimum calories floor, got %d", info.TargetCalories)
	}
}

func TestServiceCalculateTargets_ReturnsError(t *testing.T) {
	service := newServiceForTest(nil, nil)

	err := service.calculateTargets([]*models.UserInfo{
		{Sex: "female", Age: 28, HeightCm: 165, WeightKg: 60, TargetWeightKg: 55},
		{Sex: "male", Age: 0, HeightCm: 180, WeightKg: 80, TargetWeightKg: 75},
	})
	if err == nil {
		t.Fatalf("expected error")
	}

	if !strings.Contains(err.Error(), "недостаточно данных") {
		t.Fatalf("unexpected error: %v", err)
	}
}
