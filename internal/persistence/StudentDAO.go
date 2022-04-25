package persistence

import . "internal/entities/student"

type StudentDAOInterface interface {
	FindAll() []Student
	Find(id int) *Student
	Exist(id int) bool
	Create(student Student) bool
	Update(student Student) bool
	Delete(id int) bool
}
