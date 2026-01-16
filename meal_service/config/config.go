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
	Redis    RedisConfig    `mapstructure:"redis"`
	Kafka    KafkaConfig    `mapstructure:"kafka"`

	MealServiceSettings MealServiceSettings `mapstructure:"meal_service"`
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

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`     // localhost:6379
	Password string `mapstructure:"password"` // обычно пусто
	DB       int    `mapstructure:"db"`
}

type KafkaConfig struct {
	Brokers               []string `mapstructure:"brokers"` // ["localhost:9092"]
	MealConsumedTopicName string   `mapstructure:"meal_consumed_topic_name"`
}

type MealServiceSettings struct {
	MinNameLen     uint32 `mapstructure:"min_name_len"`
	MaxNameLen     uint32 `mapstructure:"max_name_len"`
	MaxWeightGrams uint32 `mapstructure:"max_weight_grams"`
	OFFUserAgent   string `mapstructure:"off_user_agent"`
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
