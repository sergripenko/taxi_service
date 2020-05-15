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

func (as *ApplicationsLocalStorage) GetApplication(ctx context.Context) *models.Application {
	rand.Seed(time.Now().UnixNano())
	as.mutex.Lock()
	randIndex := rand.Intn(len(as.availableApplications))
	application := as.availableApplications[randIndex]
	as.mutex.Unlock()
	as.addShowedApplication(application)
	return application
}

func (as *ApplicationsLocalStorage) GetAllApplications(ctx context.Context) []*models.Application {
	return as.showedApplications
}

// Remove one random applications and create one new
func (as *ApplicationsLocalStorage) refreshApplications() {
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(len(as.availableApplications))
	as.mutex.Lock()
	defer as.mutex.Unlock()
	// Copy last element to index i.
	as.availableApplications[randIndex] = as.availableApplications[len(as.availableApplications)-1]
	// Erase last element (write zero value).
	as.availableApplications[len(as.availableApplications)-1] = nil
	// Truncate slice.
	as.availableApplications = as.availableApplications[:len(as.availableApplications)-1]

	newApplication := &models.Application{
		Key:   services.GetRandomString(),
		Count: 0,
	}
	as.availableApplications = append(as.availableApplications, newApplication)
}

func (as *ApplicationsLocalStorage) addShowedApplication(application *models.Application) {
	//for first application
	as.mutex.Lock()
	defer as.mutex.Unlock()

	if len(as.showedApplications) == 0 {
		application.Count = 1
		as.showedApplications = append(as.showedApplications, application)

	} else {
		for _, app := range as.showedApplications {

			if application.Key == app.Key {
				var mux sync.Mutex
				mux.Lock()
				app.Count = app.Count + 1
				mux.Unlock()
				return
			}
		}
		application.Count = 1
		as.showedApplications = append(as.showedApplications, application)
	}
}
