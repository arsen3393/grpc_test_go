package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	DB   DatabaseConfig `env:"db" env-required:"true"`
	GRPC GrpcConfig     `env:"grpc" env-required:"true"`
}

type DatabaseConfig struct {
	Host               string `env:"DB_HOST" env-default:"localhost"`
	Port               string `env:"DB_PORT" env-default:"5432"`
	Name               string `env:"DB_NAME" env-required:"true"`
	User               string `env:"DB_USER" env-required:"true"`
	Password           string `env:"DB_PASSWORD" env-required:"true"`
	TimeZone           string `env:"DB_TIME_ZONE" env-default:"UTC" comment:"Часовой пояс базы данных"`
	MaxIdleConnections int    `env:"DB_MAX_IDLE_CONNECTIONS" envDefault:"40" comment:"Максимальное число простых соединений"`
	MaxOpenConnections int    `env:"DB_MAX_OPEN_CONNECTIONS" envDefault:"40" comment:"Максимальное число открытых соединений"`
}

type GrpcConfig struct {
	Port    int           `env:"GRPC_PORT" env-default:"8080"`
	Timeout time.Duration `env:"GRPC_TIMEOUT" env-default:"40s"`
	Address string        `env:"GRPC_ADDRESS" env-required:"true"`
}

func MustLoadConfig() *Config {
	configPath := ".env"
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file not found")
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config:" + err.Error())
	}

	return &cfg
}
