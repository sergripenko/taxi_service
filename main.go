package main

import (
	"taxi_service/config"
	"taxi_service/server"

	"github.com/labstack/gommon/log"

	"github.com/spf13/viper"
)

func main() {
	// init configurations
	if err := config.Init(); err != nil {
		log.Fatal("%s", err.Error())
	}
	app := server.NewApp()
	// start server
	if err := app.Run(viper.GetString("port")); err != nil {
		log.Fatal(err)
	}
}
