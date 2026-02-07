package offutils

import (
	"encoding/json"
	"strconv"
)

type OFFSearchResponse struct {
	Products []OFFProduct `json:"products"`
}

type OFFProduct struct {
	ProductName string        `json:"product_name"`
	Nutriments  OFFNutriments `json:"nutriments"`
}

type OFFNutriments struct {
	EnergyKcal    float32 `json:"energy-kcal_100g"`
	Proteins      float32 `json:"proteins_100g"`
	Fat           float32 `json:"fat_100g"`
	Carbohydrates float32 `json:"carbohydrates_100g"`
}

func (n *OFFNutriments) UnmarshalJSON(data []byte) error {
	type Alias OFFNutriments
	aux := &struct {
		EnergyKcal    interface{} `json:"energy-kcal_100g"`
		Proteins      interface{} `json:"proteins_100g"`
		Fat           interface{} `json:"fat_100g"`
		Carbohydrates interface{} `json:"carbohydrates_100g"`
		*Alias
	}{
		Alias: (*Alias)(n),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	n.EnergyKcal = parseFloat32(aux.EnergyKcal)
	n.Proteins = parseFloat32(aux.Proteins)
	n.Fat = parseFloat32(aux.Fat)
	n.Carbohydrates = parseFloat32(aux.Carbohydrates)

	return nil
}

func parseFloat32(value interface{}) float32 {
	switch v := value.(type) {
	case float64:
		return float32(v)
	case float32:
		return v
	case int:
		return float32(v)
	case string:
		if f, err := strconv.ParseFloat(v, 32); err == nil {
			return float32(f)
		}
		return 0
	default:
		return 0
	}
}
