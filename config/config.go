package config

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func InitConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("Не удалось загрузить .env файл, используются переменные окружения")
	}

	return &Config{
		SUPABASE_URL: os.Getenv("SUPABASE_URL"),
		SUPABASE_KEY: os.Getenv("SUPABASE_KEY"),
	}, nil
}
