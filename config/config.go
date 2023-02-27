package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port       string `env:"PORT"`
	DSN        string `env:"DSN"`
	DriverName string `env:"DRIVER_NAME"`
}

func GetConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}
	return &Config{
		Port:       os.Getenv("PORT"),
		DSN:        os.Getenv("DSN"),
		DriverName: os.Getenv("DRIVER_NAME"),
	}, nil

}
