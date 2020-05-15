package usecase

import (
	"context"
	"taxi_service/applications"
	"taxi_service/models"
)

type ApplicationsUseCase struct {
	applicationsRepo applications.Repository
}

func NewApplicationsUseCase(applicationsRepo applications.Repository) *ApplicationsUseCase {
	return &ApplicationsUseCase{
		applicationsRepo: applicationsRepo,
	}
}

func (a *ApplicationsUseCase) GetApplication(ctx context.Context) string {
	application := a.applicationsRepo.GetApplication(ctx)
	return application.Key
}

func (a *ApplicationsUseCase) GetAllApplications(ctx context.Context) []*models.Application {
	return a.applicationsRepo.GetAllApplications(ctx)
}
