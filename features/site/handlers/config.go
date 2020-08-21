package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/lucasasoaresmar/features-go/adapters/helpers"
	"github.com/lucasasoaresmar/features-go/features/site/models"
)

// ConfigRepository contract
type ConfigRepository interface {
	Get() (config models.Config, err error)
	Update(configChanges *models.Config) (updatedConfig models.Config, err error)
}

// Config http handlers
type Config struct {
	Repository ConfigRepository
}

// Get http handler
func (c *Config) Get(w http.ResponseWriter, req *http.Request) {
	config, err := c.Repository.Get()
	if err != nil {
		helpers.ErrorResponse(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	helpers.SuccessReponse(w, &config, http.StatusOK)
}

// Edit http handlers
func (c *Config) Edit(w http.ResponseWriter, req *http.Request) {
	var configChanges models.Config
	err := json.NewDecoder(req.Body).Decode(&configChanges)

	if err != nil {
		helpers.ErrorResponse(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	updatedConfig, err := c.Repository.Update(&configChanges)
	if err != nil {
		helpers.ErrorResponse(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	helpers.SuccessReponse(w, &updatedConfig, http.StatusOK)
}
