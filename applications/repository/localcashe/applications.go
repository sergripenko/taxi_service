package localcashe

import (
	"context"
	"math/rand"
	"sync"
	"taxi_service/models"
	"taxi_service/services"
	"time"
)

type ApplicationsLocalStorage struct {
	mutex                 *sync.Mutex
	availableApplications []*models.Application
	allApplications       []*models.Application
}

func NewApplicationsLocalStorage(limit int) *ApplicationsLocalStorage {
	appStorage := &ApplicationsLocalStorage{
		mutex:                 new(sync.Mutex),
		availableApplications: []*models.Application{},
		allApplications:       []*models.Application{},
	}

	for i := 0; i < limit; i++ {
		tempApp := models.Application{
			Key: services.GetRandomString(),
		}
		appStorage.availableApplications = append(appStorage.availableApplications, &tempApp)
		appStorage.allApplications = append(appStorage.allApplications, &tempApp)
	}

	ticker := time.NewTicker(time.Millisecond * 200)
	go func() {
		for range ticker.C {
			appStorage.refreshAvailableAppPool()
		}
	}()
	return appStorage
}

func (a *ApplicationsLocalStorage) GetApplication(ctx context.Context) *models.Application {
	rand.Seed(time.Now().UnixNano())
	a.mutex.Lock()
	randIndex := rand.Intn(len(a.availableApplications))
	application := a.availableApplications[randIndex]
	a.mutex.Unlock()
	a.incrementCount(application)
	return application
}

func (a *ApplicationsLocalStorage) GetAllApplications(ctx context.Context) []*models.Application {
	var applications []*models.Application
	a.mutex.Lock()

	//add app with count greater then 0
	for _, app := range a.allApplications {

		if app.Count != 0 {
			applications = append(applications, app)
		}
	}
	a.mutex.Unlock()
	return applications
}

// Remove one random applications and create one new
func (a *ApplicationsLocalStorage) refreshAvailableAppPool() {
	rand.Seed(time.Now().UnixNano())
	a.mutex.Lock()
	randIndex := rand.Intn(len(a.availableApplications))

	// Copy last element to index i.
	a.availableApplications[randIndex] = a.availableApplications[len(a.availableApplications)-1]
	// Erase last element (write zero value).
	a.availableApplications[len(a.availableApplications)-1] = nil
	// Truncate slice.
	a.availableApplications = a.availableApplications[:len(a.availableApplications)-1]

	newApplication := &models.Application{
		Key: services.GetRandomString(),
	}
	a.availableApplications = append(a.availableApplications, newApplication)
	a.allApplications = append(a.allApplications, newApplication)
	a.mutex.Unlock()
}

func (a *ApplicationsLocalStorage) incrementCount(application *models.Application) {
	//for first application
	a.mutex.Lock()
	defer a.mutex.Unlock()

	for _, app := range a.allApplications {

		if application.Key == app.Key {
			app.Count = app.Count + 1
			return
		}
	}
}
