package handlers

import (
	"ShoeStore/database/queries"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetUsers(c *gin.Context) {
	db, exists := c.Get("db")
	if !exists {
		logrus.Errorln("Не удалось получить подключение к БД")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}
	conn, ok := db.(*pgx.Conn)
	if !ok {
		logrus.Errorln("Ошибка приведения типа подключения")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}

	sneakers, err := queries.GetAllSneakers(conn)
	if err != nil {
		logrus.Errorln("Ошибка получения объектов таблицы sneakers: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения данных"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"sneakers": sneakers})
}
