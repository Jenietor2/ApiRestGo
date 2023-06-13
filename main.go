package main

import (
	"apirest/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//rutas
	router := mux.NewRouter().StrictSlash(true)

	//endpoint

	router.HandleFunc("/", handlers.IndexRoute)
	router.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	router.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	router.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
