package bootstrap

import (
	useroutbox "github.com/JoePeach762/PP_project/user_service/internal/outbox"
	userproducer "github.com/JoePeach762/PP_project/user_service/internal/producer/user"
	userstorage "github.com/JoePeach762/PP_project/user_service/internal/storage/pgstorage/userstorage"
)

func InitOutboxPublisher(
	storage *userstorage.PGstorage,
	producer *userproducer.UserKafkaProducer,
) *useroutbox.Publisher {
	return useroutbox.NewPublisher(storage, producer)
}
