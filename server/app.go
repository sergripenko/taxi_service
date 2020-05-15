package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"taxi_service/applications"
	apphttp "taxi_service/applications/delivery/http"
	"taxi_service/applications/repository/localcashe"
	"taxi_service/applications/usecase"
	"time"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

type App struct {
	httpServer     *http.Server
	applicationsUC applications.UseCase
}

func NewApp() *App {
	applicationsRepo := localcashe.NewApplicationsLocalStorage(viper.GetInt("applications_limit"))
	return &App{
		applicationsUC: usecase.NewApplicationsUseCase(applicationsRepo),
	}
}

// Run server
func (a *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	// Set up http handlers
	apphttp.RegisterHTTPEndpoints(router, a.applicationsUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return a.httpServer.Shutdown(ctx)
}
