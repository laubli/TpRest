package web

import (
	"encoding/json"
	. "entities"
	"fmt"
	"net/http"
	. "persistence"

	"github.com/gorilla/mux"
)

var (
	languages []Language // slice (dynamically sized array)
)

func init() {
	languages = []Language{Language{
		Code: "1",
		Name: "Nom 1",
	}}
}

func GetLanguage(res http.ResponseWriter, req *http.Request) {
	// We generally interact with api's in JSON
	res.Header().Set("Content-type", "application/json")

	languageCode := mux.Vars(req)["Id"]

	var languageFind = FindLanguage(languageCode)
	// returns the json encoding of posts
	result, err := json.Marshal(languageFind)

	// check for error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError) // status: 500
		res.Write([]byte(`{"error": "Error marshalling the langauge array"}`))
		return
	}

	res.WriteHeader(http.StatusOK) // status: 200
	res.Write(result)
}

func AddLanguage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	// declaring new post of type Post
	var language Language

	// reads the JSON value and decodes it into a Go value
	err := json.NewDecoder(req.Body).Decode(&language)

	// check for error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	if !CreateLanguage(language) {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error language already exist"}`))
		return
	}

	// returns the json encoding of post
	result, err := json.Marshal(language)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func CreateOneLanguage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<html>")
	fmt.Fprintf(res, "<head>")
	fmt.Fprintf(res, "</head>")
	fmt.Fprintf(res, "<body>")
	fmt.Fprintf(res, "<h1> Création d'un langague </h1>")
	fmt.Fprintf(res, "<div><label for=\"Code\">Code :</label><input type=\"text\" id=\"codeInput\" value=\" \"> </div>")
	fmt.Fprintf(res, "<div><label for=\"LastName\">Name :</label><input type=\"text\" id=\"nameInput\" value=\" \">div>")
	fmt.Fprintf(res, "<div><input type=\"button\" value=\"Submit\"></div>")
	fmt.Fprintf(res, "</body>")
	fmt.Fprintf(res, "</html>")
}

func GetAllLanguage(res http.ResponseWriter, req *http.Request) {
	// We generally interact with api's in JSON
	res.Header().Set("Content-type", "application/json")

	var languageFind = FindAllLanguages()
	// returns the json encoding of posts
	result, err := json.Marshal(languageFind)

	// check for error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError) // status: 500
		res.Write([]byte(`{"error": "Error marshalling the language array"}`))
		return
	}

	res.WriteHeader(http.StatusOK) // status: 200
	res.Write(result)
}
