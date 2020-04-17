package main

import (
	"taxi_service/config"
	"taxi_service/server"

	"github.com/labstack/gommon/log"

	appl "taxi_service/applications/delivery/http"

	"github.com/spf13/viper"
)

func main() {
	// init configurations
	if err := config.Init(); err != nil {
		log.Fatal("%s", err.Error())
	}

	//	generate applications
	appl.GenApplications(viper.GetInt("applications_limit"))

	// start server
	if err := server.Run(viper.GetString("port")); err != nil {
		log.Fatal(err)
	}
}
