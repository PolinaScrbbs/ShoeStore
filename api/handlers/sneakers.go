package handlers

import (
	"ShoeStore/database/queries"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetSneakers(c *gin.Context) {
	logrus.Infoln("Запрос на получение списка кроссовок")

	db, exists := c.Get("db")
	if !exists {
		logrus.Errorln("Ошибка: База данных недоступна")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "База данных недоступна"})
		return
	}

	pool, ok := db.(*pgxpool.Pool)
	if !ok {
		logrus.Errorln("Ошибка: Невозможно привести объект к *pgxpool.Pool")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка преобразования пула соединений"})
		return
	}

	logrus.Infoln("Подключение к базе данных успешно получено")

	sneakers, err := queries.GetAllSneakers(pool)
	if err != nil {
		logrus.Errorf("Ошибка получения объектов таблицы sneakers: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"})
		return
	}

	logrus.Infof("Успешно получено %d кроссовок", len(sneakers))
	c.JSON(http.StatusOK, gin.H{"sneakers": sneakers})
}
