package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	SecretKey  string `mapstructure:"SECRET_KEY"`
	Debug      bool   `mapstructure:"DEBUG"`
	PgUser     string `mapstructure:"POSTGRES_USER"`
	PgPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PgHost     string `mapstructure:"POSTGRES_HOST"`
	PgPort     int    `mapstructure:"POSTGRES_PORT"`
	PgDB       string `mapstructure:"POSTGRES_DB"`
	LogLevel   int    `mapstructure:"LOG_LEVEL"`
	DbIdleCon  int    `mapstructure:"DB_POOL_IDLE"`
	DbMaxCon   int    `mapstructure:"DB_POOL_MAX"`
	DbLifeTime int    `mapstructure:"DB_POOL_LIFETIME"`
	DbIdleTime int    `mapstructure:"DB_POOL_IDLETIME"`
}

func NewConfig(path string) (config *Config) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading config file %w", err))
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("error mapping config %w", err))
	}

	return
}
