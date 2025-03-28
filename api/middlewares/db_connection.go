package middlewares

import (
	"ShoeStore/database"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func DBMiddleware(url *string) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn := database.NewDBConnection(*url)
		defer func() {
			err := conn.Close(context.Background())
			if err != nil {
				logrus.Errorf("Ошибка закрытия подключения к базе данных: %v", err)
			}
		}()

		c.Set("db", conn)
		c.Next()
	}
}
