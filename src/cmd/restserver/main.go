package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter() // instance of mux router
	const port string = ":8000"

	// base route of the API
	//router.HandleFunc("/", )

	router.HandleFunc("/students", web.GetPosts).Methods("GET")
	router.HandleFunc("/students", web.AddPost).Methods("POST")

	// logging in terminal
	log.Println("Server Listening on port", port)
	log.Fatal(http.ListenAndServe(port, router)) // listening to server
}
