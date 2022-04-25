package entities

import "fmt"

type Language struct {
	Code string
	Name string
}

func NewLanguage() Language {
	return Language{"0", "name"}
}

func (language Language) String() string {
	return fmt.Sprint(language.Code) + " " + language.Name
}
