package web

import (
	"encoding/json"
	"fmt"
	entities "internal/entities/language"
	"internal/persistence"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// In case you want to use the memory implementation, use this line of code instead of the next one
var intLang persistence.LanguageDAOMemory

//var intLang persistence.LanguageDAOBolt
var lang entities.Language

func GetOneLanguage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("je suis 1")
	selectedLanguage := intLang.Find(mux.Vars(r)["code"])

	w.Header().Set("Content-Type", "application/json")

	if selectedLanguage == nil {
		w.WriteHeader(http.StatusNotFound)
			fmt.Println("je suis 2")
		return
	}

	resp, err := json.Marshal(selectedLanguage)

	// Error handling
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Error: %s", err)
			fmt.Println("je suis 3")
	}

	w.Write(resp)
}

func GetAllLanguage(w http.ResponseWriter, r *http.Request) {
	// Header
	w.Header().Set("Content-Type", "application/json")

	// Body
	resp, err := json.Marshal(intLang.FindAll())

	// Error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error happened in JSON marshal. Error: %s", err)
	}

	// Send respond
	w.WriteHeader(http.StatusFound)
	w.Write(resp)

	return
}

func CreateLanguageHandler(w http.ResponseWriter, r *http.Request) {
	// Header
	w.Header().Set("Content-Type", "application/json")

	// Parsing json
	reqBody, err := ioutil.ReadAll(r.Body)

	// Error handling parsing r.Body
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, "Error: cannot read body: %s", err)
		return
	}

	err = json.Unmarshal(reqBody, &lang)

	// Error handling when unmarshaling json
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, "Error: cannot unmarshal json: %s", err)
		return
	}

	// Body
	if !intLang.Create(lang) {
		// Language already exist
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, "Language already exist")
		return
	}

	// Language successfully created
	w.WriteHeader(http.StatusCreated)
}

func UpdateLanguageHandler(w http.ResponseWriter, r *http.Request) {
	// Header
	w.Header().Set("Content-Type", "application/json")

	// Parsing json
	reqBody, err := ioutil.ReadAll(r.Body)

	// Error handling parsing r.Body
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, "Error: cannot read body")
		return
	}

	err = json.Unmarshal(reqBody, &lang)

	// Error handling when unmarshaling json
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, "Error: cannot unmarshal json")
		return
	}

	// Update the language
	if intLang.Update(lang) {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Success: language modified")
		return
	}

	// Cannot be modified: Language cannot exist
	w.WriteHeader(http.StatusNotModified)
	fmt.Fprintf(w, "Error: language doesn't exist")
}

func DeleteLanguageByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Try to delete
	if intLang.Delete(mux.Vars(r)["code"]) {
		// Status success
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Success: language deleted")
		return
	}

	// Status error
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Error: language doesn't exist")
}
