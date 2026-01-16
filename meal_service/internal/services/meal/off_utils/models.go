package offutils

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
