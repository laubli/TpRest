package entities

type Student struct {
	Id           int
	FirstName    string
	LastName     string
	Age          int
	LanguageCode string
}

type StudentList []Student

func NewStudent(id int, firstName string, lastName string, age int, languageCode string) Student {
	return Student{
		Id:           id,
		FirstName:    firstName,
		LastName:     lastName,
		Age:          age,
		LanguageCode: languageCode,
	}
}

func (l StudentList) Len() int {
	return len(l)
}

func (l StudentList) Less(i, j int) bool {
	return l[i].Id < l[j].Id
}

func (l StudentList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func StudentString(student Student) string {
	return "{" +
		"	Etudiant number:" + string(rune(student.Id)) +
		"	First name: " + student.FirstName +
		"	Last name: " + student.LastName +
		"	Age: " + string(rune(student.Age)) +
		"	Language code: " + student.LanguageCode +
		"}"
}
