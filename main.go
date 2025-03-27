package main

import (
	"ShoeStore/config"
	"ShoeStore/database"
	"github.com/sirupsen/logrus"
)

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

	database.InitDB(conf.DatabaseUrl)
}
