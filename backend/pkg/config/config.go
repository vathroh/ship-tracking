package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort       string `mapstructure:"APP_PORT"`
	LogLevel      string `mapstructure:"LOG_LEVEL"`
	PostgresDSN   string `mapstructure:"POSTGRES_DSN"`
	RedisAddr     string `mapstructure:"REDIS_ADDR"`
	BinderByteURL string `mapstructure:"BINDERBYTE_URL"`
	BinderByteKey string `mapstructure:"BINDERBYTE_API_KEY"`
	FrontendURL   string `mapstructure:"FRONTEND_URL"`
}

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("configs")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("LOG_LEVEL", "info")
	viper.SetDefault("POSTGRES_DSN", "host=localhost user=fathur password=qwerty123 dbname=cek_ongkir port=5432 sslmode=disable")
	viper.SetDefault("REDIS_ADDR", "localhost:6379")
	viper.SetDefault("BINDERBYTE_URL", "https://api.binderbyte.com/v1")
	viper.SetDefault("FRONTEND_URL", "http://localhost:5173")

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
