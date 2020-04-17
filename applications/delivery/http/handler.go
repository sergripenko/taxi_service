package http

import (
	"math/rand"
	"net/http"
	"sync"
	"taxi_service/models"
	"taxi_service/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var AvailableApplications []*models.Application
var ShowedApplications []*models.Application

func GetApplication(c *gin.Context) {
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(viper.GetInt("applications_limit"))
	application := AvailableApplications[randIndex]
	addShowedApplication(*application)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": application.Key})
}

func GetAllApplications(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": ShowedApplications})
}

func GenApplications(limit int) {

	for i := 0; i < limit; i++ {
		tempApp := models.Application{
			Key:   services.GetRandomString(),
			Count: 0,
		}
		AvailableApplications = append(AvailableApplications, &tempApp)
	}

	ticker := time.NewTicker(time.Millisecond * 200)
	go func() {
		for range ticker.C {
			refreshApplications(AvailableApplications)
		}
	}()
}

func refreshApplications(apps []*models.Application) {
	rand.Seed(time.Now().UnixNano())
	randIndex := rand.Intn(viper.GetInt("applications_limit"))
	var mux sync.Mutex
	mux.Lock()
	// Copy last element to index i.
	apps[randIndex] = apps[len(apps)-1]
	// Erase last element (write zero value).
	apps[len(apps)-1] = nil
	// Truncate slice.
	apps = apps[:len(apps)-1]
	mux.Unlock()

	newApplication := &models.Application{
		Key:   services.GetRandomString(),
		Count: 0,
	}
	apps = append(apps, newApplication)
}

func addShowedApplication(application models.Application) {
	//for first application
	if len(ShowedApplications) == 0 {
		application.Count = 1
		ShowedApplications = append(ShowedApplications, &application)

	} else {
		for _, app := range ShowedApplications {

			if application.Key == app.Key {
				var mux sync.Mutex
				mux.Lock()
				app.Count = app.Count + 1
				mux.Unlock()
				return
			}
		}
		application.Count = 1
		ShowedApplications = append(ShowedApplications, &application)
	}
}
