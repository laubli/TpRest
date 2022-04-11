package entities

import "fmt"

type Student struct {
	Id          int
	FirstName   string
	LastName    string
	Age         int
	LangageCode bool
}

func NewStudent() Student {
	return Student{0, "Fname", "Lname", 10, true}
}

func (student Student) String() string {
	return fmt.Sprint(student.Id) + " " + student.FirstName + " " + student.LastName + " " + fmt.Sprint(student.Age) + " " + fmt.Sprint(student.LangageCode)
}
