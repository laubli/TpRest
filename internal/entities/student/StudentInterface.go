package entities

// This interface serve the purpose of being able to use the sort.Sort()
// function with the struct Student
type StudentInterface interface {
	Len() int

	Less(i, j int) bool

	Swap(i, j int)
}
