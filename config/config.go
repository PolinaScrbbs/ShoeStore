package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func InitConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("Не удалось загрузить .env файл, используются переменные окружения")
	}

	config := &Config{
		USER:     os.Getenv("USER"),
		PASSWORD: os.Getenv("PASSWORD"),
		HOST:     os.Getenv("HOST"),
		PORT:     os.Getenv("PORT"),
		DB_NAME:  os.Getenv("DB_NAME"),
	}

	config.DatabaseUrl = config.GetDataBaseUrl()

	return config, nil
}

func (c Config) GetDataBaseUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.USER, c.PASSWORD, c.HOST, c.PORT, c.DB_NAME)
}
