package persistence

import (
	"encoding/json"
	"fmt"
	. "internal/entities"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateLanguage(w http.ResponseWriter, r *http.Request) {
	var newLanguage Language
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Une erreur s'est produit")
	}

	json.Unmarshal(reqBody, &newLanguage)
	languages = append(languages, newLanguage)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newLanguage)
}

func getOneLanguage(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, singleLanguage := range languages {
		if singleLanguage.Code == languageCode {
			json.NewEncoder(w).Encode(singleLanguage)
		}
	}
}

func getAllLanguages(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(languages)
}

func updateLanguage(w http.ResponseWriter, r *http.Request) {
	languageCode := mux.Vars(r)["Code"]
	var updatedLanguage Language

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Une erreur s'est produit")
	}
	json.Unmarshal(reqBody, &updatedLanguage)

	for i, singleLanguage := range languages {
		if singleLanguage.Code == languageCode {
			singleLanguage.Name = updatedLanguage.Name
			languages = append(languages[:i], singleLanguage)
			json.NewEncoder(w).Encode(singleLanguage)
		}
	}
}

func deleteLanguage(w http.ResponseWriter, r *http.Request) {
	LanguageCode := mux.Vars(r)["Code"]

	for i, singleLanguage := range languages {
		if singleLanguage.Code == LanguageCode {
			languages = append(languages[:i], languages[i+1:]...)
			fmt.Fprintf(w, "The langauge with code %v has been deleted successfully", LanguageCode)
		}
	}
}
