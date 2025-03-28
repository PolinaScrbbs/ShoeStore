package database

import (
	"ShoeStore/database/queries/creation"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"time"
)

var DB *pgxpool.Pool

func InitDB(url *string) {
	pool, err := pgxpool.New(context.Background(), *url)
	if err != nil {
		logrus.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	DB = pool

	go func() {
		_, err := DB.Exec(context.Background(), creation.CreateSneakersTableSQL)
		if err != nil {
			logrus.Fatalf("Ошибка создания таблицы: %v", err)
		}
		logrus.Infoln("Таблица успешно создана или уже существует.")
	}()

	time.Sleep(2 * time.Second)

	logrus.Infoln("Инициализация БД успешно завершена")
}
