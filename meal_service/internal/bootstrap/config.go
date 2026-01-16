package bootstrap

import (
	"fmt"
	"strings"

	"github.com/JoePeach762/PP_project/meal_service/config"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (*config.Config, error) {
	_ = godotenv.Load() // .env для локальной разработки

	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("APP")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("ошибка чтения конфига: %w", err)
	}

	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("ошибка парсинга конфига: %w", err)
	}

	return &cfg, nil
}
