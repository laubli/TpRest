package web

import (
	"encoding/json"
	. "entities"
	"net/http"
	. "persistence"
	"strconv"

	"github.com/gorilla/mux"
)

func AddStudent(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	// declaring new post of type Post
	var student Student

	// reads the JSON value and decodes it into a Go value
	err := json.NewDecoder(req.Body).Decode(&student)

	// check for error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	if !CreateStudent(student) {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error student already exist"}`))
		return
	}

	// returns the json encoding of post
	result, err := json.Marshal(student)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func GetStudent(res http.ResponseWriter, req *http.Request) {
	// We generally interact with api's in JSON
	res.Header().Set("Content-type", "application/json")

	studentIdString := mux.Vars(req)["Id"]
	studentId, errStr := strconv.Atoi(studentIdString)
	if errStr != nil {
		res.WriteHeader(http.StatusInternalServerError) // status: 500
		res.Write([]byte(`{"error": "Error"}`))
		return
	}

	var studentFind = FindStudent(studentId)
	// returns the json encoding of posts
	result, err := json.Marshal(studentFind)

	// check for error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError) // status: 500
		res.Write([]byte(`{"error": "Error marshalling the student array"}`))
		return
	}

	res.WriteHeader(http.StatusOK) // status: 200
	res.Write(result)
}

func GetAllStudents(res http.ResponseWriter, req *http.Request) {
	// We generally interact with api's in JSON
	res.Header().Set("Content-type", "application/json")

	var studentFind = FindAllStudents()
	// returns the json encoding of posts
	result, err := json.Marshal(studentFind)

	// check for error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError) // status: 500
		res.Write([]byte(`{"error": "Error marshalling the student array"}`))
		return
	}

	res.WriteHeader(http.StatusOK) // status: 200
	res.Write(result)
}
