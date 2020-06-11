package router

import (
	"../middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/get", middleware.GetAllEntry).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/add", middleware.CreateEntry).Methods("POST", "OPTIONS")
	return router
}
