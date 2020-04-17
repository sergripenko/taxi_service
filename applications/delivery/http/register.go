package http

import (
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.Engine) {
	router.GET("/request", GetApplication)
	router.GET("/admin/request", GetAllApplications)
}
