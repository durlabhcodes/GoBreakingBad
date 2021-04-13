package main

import (
	"GoBreakingBad/routes"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

//HandleRequests is an amazing func
func HandleRequests() {
	router := mux.NewRouter()
	setupRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Info("Server start at http://localhost:8080")
}

func setupRoutes(router *mux.Router) {
	log.Info("Setting up routes")
	router.HandleFunc("/ping", routes.Ping).Methods("GET")
}

func getApiVersion() string {
	return "v1"
}
