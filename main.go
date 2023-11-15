package main

import (
	"serviceemployee/configs"
	"serviceemployee/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.EmployeeRoutes(router)


	log.Println("Server running on port 9000")
	http.ListenAndServe(":9000", router)
}