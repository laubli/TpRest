package persistence

import (
	"encoding/json"
	"fmt"
	. "internal/entities"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	students []Student // slice (dynamically sized array)
)

func init() {
	students = []Student{entities.Student{
		Id:          1,
		FirstName:   "Nom 1",
		LastName:    "Prénom 1",
		Age:         20,
		LangageCode: true,
	}}
}

func createStudent(w http.ResponseWriter, r *http.Request) {
	var newStudent Student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Erreur")
	}

	json.Unmarshal(reqBody, &newStudent)
	students = append(students, newStudent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newStudent)
}

func getOneStudent(students []Student, studentId int) (student Student) {
	//studentId := mux.Vars(r)["Id"]
	var returnStudent Student

	for _, singleStudent := range students {
		if singleStudent.Id == studentId {
			returnStudent = singleStudent
			//json.NewEncoder(w).Encode(singleStudent)
		}
	}
	return returnStudent
}

func getAllStudents(w http.ResponseWriter, r *http.Request) (students []Student) {
	json.NewEncoder(w).Encode(students)
}

func updateStudent(w http.ResponseWriter, r *http.Request) {
	studentId := mux.Vars(r)["Id"]
	var updatedStudent Student

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "erreur")
	}
	json.Unmarshal(reqBody, &updatedStudent)

	for i, singleStudent := range students {
		if singleStudent.Id == studentId {
			singleStudent.FirstName = updatedStudent.FirstName
			singleStudent.LastName = updatedStudent.LastName
			singleStudent.Age = updatedStudent.Age
			singleStudent.LangageCode = updatedStudent.LangageCode
			students = append(students[:i], singleStudent)
			json.NewEncoder(w).Encode(singleStudent)
		}
	}
}

func deleteStudent(w http.ResponseWriter, r *http.Request) {
	studentId := mux.Vars(r)["id"]

	for i, singleStudent := range evestudentsnts {
		if singleStudent.Id == studentId {
			students = append(students[:i], students[i+1:]...)
			fmt.Fprintf(w, "The student with Id %v has been deleted successfully", studentId)
		}
	}
}
