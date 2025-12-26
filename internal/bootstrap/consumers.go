package bootstrap

import (
	"fmt"

	"github.com/JoePeach762/PP_project/config"
	studentsinfoupsertconsumer "github.com/JoePeach762/PP_project/internal/consumer/students_Info_upsert_consumer"
	studentsinfoprocessor "github.com/JoePeach762/PP_project/internal/services/processors/students_info_processor"
)

func InitStudentInfoUpsertConsumer(cfg *config.Config, studentsInfoProcessor *studentsinfoprocessor.StudentsInfoProcessor) *studentsinfoupsertconsumer.StudentInfoUpsertConsumer {
	kafkaBrockers := []string{fmt.Sprintf("%v:%v", cfg.Kafka.Host, cfg.Kafka.Port)}
	return studentsinfoupsertconsumer.NewStudentInfoUpsertConsumer(studentsInfoProcessor, kafkaBrockers, cfg.Kafka.StudentInfoUpsertTopicName)
}
