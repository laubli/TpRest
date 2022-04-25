package entities

// This interface serve the purpose of being able to use the sort.Sort()
// function with the struct Language
type LanguageInterface interface {
	Len() int

	Less(i, j int) bool

	Swap(i, j int)
}
