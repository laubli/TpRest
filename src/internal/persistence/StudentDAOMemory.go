package persistence

import (
	. "internal/entities/student"
	"sort"
)

type StudentDAOMemory struct{}

var _ StudentDAOInterface = (*StudentDAOMemory)(nil)

var (
	students []Student // slice (dynamically sized array)
)

func init() {
	students = []Student{Student{
		Id:          1,
		FirstName:   "Nom 1",
		LastName:    "Prénom 1",
		Age:         20,
		LanguageCode: "true",
	}}
}

func (s StudentDAOMemory) Exist(testStudent int) bool {
	for _, singleStudent := range students {
		if singleStudent.Id == testStudent {
			return true
		}
	}
	return false
}

func (s StudentDAOMemory) Create(newStudent Student) bool {
	if !s.Exist(newStudent.Id) {
		students = append(students, newStudent)
		return true
	}
	return false
}

func (s StudentDAOMemory) Find(studentId int) *Student {
	var returnStudent Student

	for _, singleStudent := range students {
		if singleStudent.Id == studentId {
			returnStudent = singleStudent
		}
	}
	return &returnStudent
}

func (s StudentDAOMemory) FindAll() []Student {
	sort.Sort(StudentList(students))
	return students
}

func (s StudentDAOMemory) Update(updatedStudent Student)  bool {
	for i, singleStudent := range students {
		if singleStudent.Id == updatedStudent.Id {
			singleStudent.FirstName = updatedStudent.FirstName
			singleStudent.LastName = updatedStudent.LastName
			singleStudent.Age = updatedStudent.Age
			singleStudent.LanguageCode = updatedStudent.LanguageCode
			students = append(students[:i], singleStudent)
			return true
		}
	}
	return false
}

func (s StudentDAOMemory) Delete(studentId int) bool {
	for i, singleStudent := range students {
		if singleStudent.Id == studentId {
			students = append(students[:i], students[i+1:]...)
			return true
		}
	}
	return false
}
