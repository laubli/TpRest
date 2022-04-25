package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/laubli/TpRest/src/internal/web"
)

func main() {
	router := mux.NewRouter() // instance of mux router
	const port string = ":8000"

	// base route of the API
	//router.HandleFunc("/", )

	router.HandleFunc("/students", web.GetStudent).Methods("GET")
	router.HandleFunc("/students", web.AddStudent).Methods("POST")
	router.HandleFunc("/languages", web.GetLanguage).Methods("GET")
	router.HandleFunc("/languages", web.AddLanguage).Methods("POST")
	router.HandleFunc("/language", web.CreateLanguage)

	// logging in terminal
	log.Println("Server Listening on port", port)
	log.Fatal(http.ListenAndServe(port, router)) // listening to server
}
