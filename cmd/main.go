package main

import (
	"ShoeStore/api/middlewares"
	"ShoeStore/api/routers"
	"ShoeStore/config"
	"ShoeStore/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "ShoeStore/docs"
)

// @title           Sneakers Store API
// @version         1.0
// @description     The sneakers store REST-API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})
	logrus.SetLevel(logrus.InfoLevel)

	conf, err := config.InitConfig()
	if err != nil {
		logrus.Error(err)
	}

	database.InitDB(&conf.DatabaseUrl)

	r := gin.Default()
	r.Use(middlewares.DBMiddleware(database.DB))

	r = routers.SneakersRouter(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(":8080"); err != nil {
		logrus.Fatal("Ошибка запуска сервера: ", err)
	}
}
