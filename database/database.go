package database

import (
	"ShoeStore/database/queries/creation"
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"time"
)

func NewDBConnection(url string) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		logrus.Fatalf("Не удалось подключиться к базе данных: %w", err)
	}
	return conn
}

func CreateTable(conn *pgx.Conn, tableName string, query string) {
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		logrus.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	logrus.Infof("Таблица %s успешно создана или уже существует.", tableName)
}

func InitDB(url *string) {
	conn := NewDBConnection(*url)
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			logrus.Fatalf("Ошибка закрытия подключения к базе данных: %v", err)
		}
	}(conn, context.Background())

	go CreateTable(conn, creation.TableName, creation.CreateSneakersTableSQL)

	time.Sleep(5 * time.Second)

	logrus.Infoln("Инициализация БД успешна завершена")
}
