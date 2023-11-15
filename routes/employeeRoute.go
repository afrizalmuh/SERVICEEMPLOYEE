package routes

import (
	"github.com/gorilla/mux"
	"serviceemployee/controllers"
	"serviceemployee/middlewares"
)

func EmployeeRoutes(r *mux.Router){
	router := r.PathPrefix("/employee").Subrouter()

	router.Use(middlewares.Auth)

	router.HandleFunc("/me", controllers.GetEmployee).Methods("GET")
	router.HandleFunc("/me/{id}", controllers.UpdateEmployee).Methods("PUT")
}