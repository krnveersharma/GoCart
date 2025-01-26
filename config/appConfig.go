package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	ServerPort string
	Dsn        string
}

func SetupEnv() (cfg AppConfig, err error) {
	godotenv.Load()
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		return AppConfig{}, errors.New("http port is not found")
	}

	Dsn := os.Getenv("DSN")
	if httpPort == "" {
		return AppConfig{}, errors.New("DSN is not found")
	}
	return AppConfig{ServerPort: httpPort, Dsn: Dsn}, nil
}
