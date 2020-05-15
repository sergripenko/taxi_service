package http

import (
	"net/http"
	"taxi_service/applications"
	"taxi_service/models"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase applications.UseCase
}

func NewHandler(useCase applications.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

type getApplication struct {
	Application string `json:"application"`
}

func (h *Handler) GetApplication(c *gin.Context) {
	appKey := h.useCase.GetApplication(c.Request.Context())
	c.JSON(http.StatusOK, &getApplication{
		Application: appKey,
	})
}

type getAllApplications struct {
	Applications []*models.Application `json:"applications"`
}

func (h *Handler) GetAllApplications(c *gin.Context) {
	apps := h.useCase.GetAllApplications(c.Request.Context())
	c.JSON(http.StatusOK, &getAllApplications{
		Applications: apps,
	})
}
