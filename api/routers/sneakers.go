package routers

import (
	"ShoeStore/api/handlers"
	"ShoeStore/api/middlewares"
	"github.com/gin-gonic/gin"
)

func SneakersRouter(url *string) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.DBMiddleware(url))

	r.GET("/sneakers", handlers.GetUsers)
	return r
}
