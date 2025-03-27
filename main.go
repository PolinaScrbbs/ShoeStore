package main

import (
	"ShoeStore/config"
	"github.com/sirupsen/logrus"
	"github.com/supabase-community/supabase-go"
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

	logrus.Infof("SUPABASE_URL: %s", conf.SUPABASE_URL)
	logrus.Infof("SUPABASE_KEY: %.15s...", conf.SUPABASE_KEY)

	_, err = supabase.NewClient(conf.SUPABASE_URL, conf.SUPABASE_KEY, &supabase.ClientOptions{})
	if err != nil {
		logrus.Fatalf("Cannot initalize client %s", err)
	}
	logrus.Infoln("Supabase is enabled")
}
