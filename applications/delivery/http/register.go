package http

import (
	"taxi_service/applications"

	"github.com/gin-gonic/gin"
)

// RegisterHTTPEndpoints - registering routes
func RegisterHTTPEndpoints(router *gin.Engine, au applications.UseCase) {
	h := NewHandler(au)
	router.GET("/request", h.GetApplication)
	router.GET("/admin/request", h.GetAllApplications)
}
