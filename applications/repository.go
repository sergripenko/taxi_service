package applications

import (
	"context"
	"taxi_service/models"
)

type Repository interface {
	GetApplication(ctx context.Context) *models.Application
	GetAllApplications(ctx context.Context) []*models.Application
}
