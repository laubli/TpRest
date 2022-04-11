package web

import (
	"encoding/json"
	"internal/entities"
	"net/http"
)

var (
	students []entities.Student // slice (dynamically sized array)
)

func init() {
	students = []entities.Student{entities.Student{
		Id:          1,
		FirstName:   "Nom 1",
		LastName:    "Prénom 1",
		Age:         20,
		LangageCode: true,
	}}
}

func GetPosts(res http.ResponseWriter, req *http.Request) {

	// We generally interact with api's in JSON
	res.Header().Set("Content-type", "application/json")

	// returns the json encoding of posts
	result, err := json.Marshal(students)

	// check for error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError) // status: 500
		res.Write([]byte(`{"error": "Error marshalling the student array"}`))
		return
	}

	res.WriteHeader(http.StatusOK) // status: 200
	res.Write(result)
}

func AddPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-type", "application/json")

	// declaring new post of type Post
	var student entities.Student

	// reads the JSON value and decodes it into a Go value
	err := json.NewDecoder(req.Body).Decode(&student)

	// check for error
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}

	// fake ID for the post
	student.Id = len(students) + 1

	// appending the post at the end of dummy array
	students = append(students, student)
	res.WriteHeader(http.StatusOK)

	// returns the json encoding of post
	result, err := json.Marshal(student)
	res.Write(result)
}
