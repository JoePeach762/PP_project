package config

import (
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort int `mapstructure:"http_port"`

	Database DatabaseConfig `mapstructure:"database"`
	Kafka    KafkaConfig    `mapstructure:"kafka"`

	UserServiceSettings UserServiceSettings `mapstructure:"user_service"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

func (d *DatabaseConfig) ConnString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		d.Username, d.Password, d.Host, d.Port, d.DBName, d.SSLMode)
}

type KafkaConfig struct {
	Brokers               []string `mapstructure:"brokers"` // ["localhost:9092"]
	MealConsumedTopicName string   `mapstructure:"meal_consumed_topic_name"`
}

type UserServiceSettings struct {
	MinNameLen uint8 `mapstructure:"min_name_len"`
	MaxNameLen uint8 `mapstructure:"max_name_len"`
	MinWeight  uint8 `mapstructure:"min_weight"`
	MaxWeight  uint8 `mapstructure:"max_weight"`
}

func LoadConfig(configPath string) (*Config, error) {
	_ = godotenv.Load()

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, fmt.Errorf("config file not found: %s", configPath)
		}
		return nil, fmt.Errorf("error reading config: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode config: %w", err)
	}

	return &cfg, nil
}
