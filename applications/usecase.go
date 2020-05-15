package applications

import (
	"context"
	"taxi_service/models"
)

type UseCase interface {
	GetApplication(ctx context.Context) string
	GetAllApplications(ctx context.Context) []*models.Application
}
