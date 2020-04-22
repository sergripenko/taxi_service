package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	apphttp "taxi_service/applications/delivery/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

// Run server
func Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	// Set up http handlers
	apphttp.RegisterHTTPEndpoints(router)

	// HTTP Server
	httpServer := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return httpServer.Shutdown(ctx)
}
