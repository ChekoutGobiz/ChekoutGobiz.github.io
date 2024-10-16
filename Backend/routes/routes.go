package routes

import (
	"github.com/gorilla/mux"
	"github.com/kuyjajan/kuyjajan-backend/controllers"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	return router
}
