package site

import (
	"fmt"
	"net/http"

	"github.com/lucasasoaresmar/features-go/features/site/handlers"
	"github.com/lucasasoaresmar/features-go/features/site/models"
	"github.com/gorilla/mux"
)

// RegisterRoutes for site feature
func RegisterRoutes(router *mux.Router) {
	siteRouter := router.PathPrefix("/site").Subrouter()

	cr := models.ConfigRepository{}
	ch := handlers.Config{Repository: &cr}

	cr.Migrate()

	router.HandleFunc("/site", ch.Get).Methods("GET")
	router.HandleFunc("/site", ch.Edit).Methods("PUT")

	siteRouter.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "pong")
	})
}
