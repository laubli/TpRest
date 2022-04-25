package entities

import "strings"

type Language struct {
	Code string
	Name string
}

type LanguageList []Language

func NewLanguage(code string, name string) Language {
	return Language{
		Code: code,
		Name: name,
	}
}

func (l LanguageList) Len() int {
	return len(l)
}

func (l LanguageList) Less(i, j int) bool {
	return strings.Compare(l[i].Code, l[j].Code) < 0
}

func (l LanguageList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func LanguageString(language Language) string {
	return "{" +
		"	Code: " + language.Code +
		"	Name: " + language.Name +
		"}"
}
