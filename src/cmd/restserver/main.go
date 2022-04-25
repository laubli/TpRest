package main

import (
	"fmt"
	. "internal/persistence"
	. "internal/web/rest"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// @title           API Go
// @version         1.0
// @description     This API brings CRUD operations about students and languages
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	fmt.Println("Meilleur API au monde, réalisé avec Gorilla Mux Routers !")
	InitialiseDB()
	defer CloseDatabase()

	// Instance of mux router
	router := mux.NewRouter()

	// Define routes and handlers
	// Languages
	router.HandleFunc("/students", GetAllStudent).Methods("GET")
	router.HandleFunc("/students/{id}", GetOneStudent).Methods("GET")
	router.HandleFunc("/students", CreateStudentHandler).Methods("POST")
	router.HandleFunc("/students", UpdateStudentHandler).Methods("PUT")
	router.HandleFunc("/students/{id}", DeleteStudentByIdHandler).Methods("DELETE")

	router.HandleFunc("/languages", GetAllLanguage).Methods("GET")
	router.HandleFunc("/languages", CreateLanguageHandler).Methods("POST")
	router.HandleFunc("languages/{code}", GetOneLanguage).Methods("GET")
	router.HandleFunc("/languages", UpdateLanguageHandler).Methods("PUT")
	router.HandleFunc("/languages/{code}", DeleteLanguageByIdHandler).Methods("DELETE")

	http.Handle("/", router)

	// Error handling
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
