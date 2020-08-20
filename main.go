package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lucasasoaresmar/features-go/features/site"
	"github.com/gorilla/mux"
)

var (
	port string = envOrDefault("PORT", ":8000")
)

func main() {
	router := mux.NewRouter()
	router.Use(defaultHeaderMiddleware)

	router.HandleFunc("/ping", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	v1Router := router.PathPrefix("/api/v1").Subrouter()
	site.RegisterRoutes(v1Router)

	log.Fatal(http.ListenAndServe(port, router))
}

func defaultHeaderMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		// w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func envOrDefault(envName string, defaultValue string) string {
	env, ok := os.LookupEnv("PORT")
	if ok {
		return env
	}
	return defaultValue
}
