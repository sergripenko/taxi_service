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
	availableApplications []*models.Application
	showedApplications    []*models.Application
	mutex                 *sync.Mutex
}

func NewApplicationsLocalStorage(limit int) *ApplicationsLocalStorage {
	appStorage := &ApplicationsLocalStorage{
		showedApplications: []*models.Application{},
		mutex:              new(sync.Mutex),
	}

	for i := 0; i < limit; i++ {
		tempApp := models.Application{
			Key:   services.GetRandomString(),
			Count: 0,
		}
		appStorage.availableApplications = append(appStorage.availableApplications, &tempApp)
	}

	ticker := time.NewTicker(time.Millisecond * 200)
	go func() {
		for range ticker.C {
			appStorage.refreshApplications()
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
	a.addShowedApplication(application)
	return application
}

func (a *ApplicationsLocalStorage) GetAllApplications(ctx context.Context) []*models.Application {
	return a.showedApplications
}

// Remove one random applications and create one new
func (a *ApplicationsLocalStorage) refreshApplications() {
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(a.availableApplications))
	a.mutex.Lock()
	defer a.mutex.Unlock()
	// Copy last element to index i.
	a.availableApplications[randIndex] = a.availableApplications[len(a.availableApplications)-1]
	// Erase last element (write zero value).
	a.availableApplications[len(a.availableApplications)-1] = nil
	// Truncate slice.
	a.availableApplications = a.availableApplications[:len(a.availableApplications)-1]

	newApplication := &models.Application{
		Key:   services.GetRandomString(),
		Count: 0,
	}
	a.availableApplications = append(a.availableApplications, newApplication)
}

func (a *ApplicationsLocalStorage) addShowedApplication(application *models.Application) {
	//for first application
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if len(a.showedApplications) == 0 {
		application.Count = 1
		a.showedApplications = append(a.showedApplications, application)

	} else {
		for _, app := range a.showedApplications {

			if application.Key == app.Key {
				var mux sync.Mutex
				mux.Lock()
				app.Count = app.Count + 1
				mux.Unlock()
				return
			}
		}
		application.Count = 1
		a.showedApplications = append(a.showedApplications, application)
	}
}
