package bootstrap

import (
	mealoutbox "github.com/JoePeach762/PP_project/meal_service/internal/outbox"
	mealproducer "github.com/JoePeach762/PP_project/meal_service/internal/producer/meal"
	mealstorage "github.com/JoePeach762/PP_project/meal_service/internal/storage/pgstorage/mealstorage"
)

func InitOutboxPublisher(
	storage *mealstorage.PGstorage,
	producer *mealproducer.MealKafkaProducer,
) *mealoutbox.Publisher {
	return mealoutbox.NewPublisher(storage, producer)
}
