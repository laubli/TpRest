package persistence

import (
	"entities"
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

func ExistStudent(testStudent entities.Student) (exist bool) {
	//reqBody, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//fmt.Fprintf(w, "Erreur")
	//}

	//json.Unmarshal(reqBody, &newStudent)
	for _, singleStudent := range students {
		if singleStudent.Id == testStudent.Id {
			return true
		}
	}
	return false
}

func CreateStudent(newStudent entities.Student) (reussi bool) {
	//reqBody, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//fmt.Fprintf(w, "Erreur")
	//}

	//json.Unmarshal(reqBody, &newStudent)
	var test = false
	for _, singleStudent := range students {
		if singleStudent.Id == newStudent.Id {
			test = true
			//json.NewEncoder(w).Encode(singleStudent)
		}
	}
	if !test {
		students = append(students, newStudent)
	}
	return test
	//w.WriteHeader(http.StatusCreated)

	//json.NewEncoder(w).Encode(newStudent)
}

func FindStudent(studentId int) (student *entities.Student) {
	//studentId := mux.Vars(r)["Id"]
	var returnStudent entities.Student

	for _, singleStudent := range students {
		if singleStudent.Id == studentId {
			returnStudent = singleStudent
			//json.NewEncoder(w).Encode(singleStudent)
		}
	}
	return &returnStudent
}

func FindAllStudents() (students []entities.Student) {
	//json.NewEncoder(w).Encode(students)
	return students
}

func UpdateStudent(updatedStudent entities.Student) (reussi bool) {
	//studentId := mux.Vars(r)["Id"]

	//reqBody, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	fmt.Fprintf(w, "erreur")
	//}
	//json.Unmarshal(reqBody, &updatedStudent)

	for i, singleStudent := range students {
		if singleStudent.Id == updatedStudent.Id {
			singleStudent.FirstName = updatedStudent.FirstName
			singleStudent.LastName = updatedStudent.LastName
			singleStudent.Age = updatedStudent.Age
			singleStudent.LangageCode = updatedStudent.LangageCode
			students = append(students[:i], singleStudent)
			//json.NewEncoder(w).Encode(singleStudent)
			return true
		}
	}
	return false
}

func DeleteStudent(studentId int) (reussi bool) {
	//studentId := mux.Vars(r)["id"]

	for i, singleStudent := range students {
		if singleStudent.Id == studentId {
			students = append(students[:i], students[i+1:]...)
			//fmt.Fprintf(w, "The student with Id %v has been deleted successfully", studentId)
			return true
		}
	}
	return false
}
