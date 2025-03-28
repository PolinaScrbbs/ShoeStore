package routers

import (
	"ShoeStore/api/handlers"
	"github.com/gin-gonic/gin"
)

// @Summary Получить все кроссовки
// @Description Возвращает список всех кроссовок в магазине
// @Tags Sneakers
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Sneaker
// @Failure 500 {object} map[string]string
// @Router /sneakers [get]
func SneakersRouter(r *gin.Engine) *gin.Engine {
	//r.Use(middlewares.DBMiddleware(url))
	r.GET("/sneakers", handlers.GetSneakers)
	return r
}
