package web

import (
	"encoding/json"
	"fmt"
	entities "internal/entities/student"
	"internal/persistence"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var intStud persistence.StudentDAOMemory
//var intStud persistence.StudentDAOBolt
var stud entities.Student

func GetOneStudent(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	w.Header().Set("Content-Type", "application/json")

	selectedStudent := intStud.Find(id)
	if selectedStudent == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resp, err := json.Marshal(selectedStudent)

	// Error handling
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Error: %s", err)
	}

	w.Write(resp)
}

func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	// Header
	w.Header().Set("Content-Type", "application/json")

	// Body
	resp, err := json.Marshal(intStud.FindAll())

	// Error handling
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Error happened in JSON marshal. Error: %s", err)
	}

	// Send respond
	w.WriteHeader(http.StatusFound)
	w.Write(resp)
	return
}

func CreateStudentHandler(w http.ResponseWriter, r *http.Request) {
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

	err = json.Unmarshal(reqBody, &stud)

	// Error handling when unmarshaling json
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, "Error: cannot unmarshal json: %s", err)
		return
	}

	// Body
	if !intStud.Create(stud) {
		// Student already exist
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, "Student already exist")
		return
	}

	// Student successfully created
	w.WriteHeader(http.StatusCreated)
}

func UpdateStudentHandler(w http.ResponseWriter, r *http.Request) {
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

	err = json.Unmarshal(reqBody, &stud)

	// Error handling when unmarshaling json
	if err != nil {
		w.WriteHeader(http.StatusNotModified)
		fmt.Fprintf(w, "Error: cannot unmarshal json: %s", err)
		return
	}

	// Update the Student
	if intStud.Update(stud) {
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Success: Student modified")
		return
	}

	// Status
	w.WriteHeader(http.StatusNotModified)
	fmt.Fprintf(w, "Error: Student doesn't exist")
}

func DeleteStudentByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Try to delete
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if intStud.Delete(id) {
		// Status success
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Success: Student deleted")
		return
	}

	// Status error
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "Error: Student doesn't exist")
}
