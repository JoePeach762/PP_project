package user

import (
	"strings"
	"testing"

	"github.com/JoePeach762/PP_project/user_service/internal/models"
)

func TestServiceValidateSingle_NormalizesSex(t *testing.T) {
	service := newServiceForTest(nil, nil)
	input := validUserInput()
	input.Sex = " FEMALE "

	if err := service.validateSingle(input); err != nil {
		t.Fatalf("validateSingle returned error: %v", err)
	}

	if input.Sex != "female" {
		t.Fatalf("expected normalized sex to be female, got %q", input.Sex)
	}
}

func TestServiceValidateSingle_InvalidFields(t *testing.T) {
	service := newServiceForTest(nil, nil)

	tests := []struct {
		name    string
		mutate  func(*models.UserInput)
		wantErr string
	}{
		{
			name: "name too short",
			mutate: func(input *models.UserInput) {
				input.Name = "A"
			},
			wantErr: "имя должно быть длиной",
		},
		{
			name: "age out of range",
			mutate: func(input *models.UserInput) {
				input.Age = 0
			},
			wantErr: "возраст должен быть",
		},
		{
			name: "height out of range",
			mutate: func(input *models.UserInput) {
				input.HeightCm = 40
			},
			wantErr: "рост должен быть",
		},
		{
			name: "weight out of range",
			mutate: func(input *models.UserInput) {
				input.WeightKg = 10
			},
			wantErr: "вес должен быть",
		},
		{
			name: "target weight out of range",
			mutate: func(input *models.UserInput) {
				input.TargetWeightKg = 10
			},
			wantErr: "целевой вес должен быть",
		},
		{
			name: "invalid sex",
			mutate: func(input *models.UserInput) {
				input.Sex = "robot"
			},
			wantErr: "некорректный пол",
		},
		{
			name: "invalid email",
			mutate: func(input *models.UserInput) {
				input.Email = "bad-email"
			},
			wantErr: "некорректный email",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := validUserInput()
			tt.mutate(input)

			err := service.validateSingle(input)
			if err == nil {
				t.Fatalf("expected error")
			}

			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("expected error to contain %q, got %q", tt.wantErr, err.Error())
			}
		})
	}
}

func TestServiceValidate_ValidatesBatch(t *testing.T) {
	service := newServiceForTest(nil, nil)
	first := validUserInput()
	second := validUserInput()
	second.Name = "Bob"
	second.Email = "bob@example.com"
	second.Sex = " MALE "

	if err := service.Validate([]*models.UserInput{first, second}); err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}

	if first.Sex != "female" {
		t.Fatalf("expected first user sex to be normalized, got %q", first.Sex)
	}

	if second.Sex != "male" {
		t.Fatalf("expected second user sex to be normalized, got %q", second.Sex)
	}
}
