package persistence

import (
	. "entities"
)

var (
	languages []Language // slice (dynamically sized array)
)

func init() {
	languages = []Language{Language{
		Code: "1",
		Name: "Nom 1",
	}}
}

func ExistLanguage(testLanguage Language) (exist bool) {
	for _, singleLanguage := range languages {
		if singleLanguage.Code == testLanguage.Code {
			return true
		}
	}
	return false
}

func CreateLanguage(newLanguage Language) (reussi bool) {
	var test = false
	for _, singleLanguage := range languages {
		if singleLanguage.Code == newLanguage.Code {
			test = true
		}
	}
	if !test {
		languages = append(languages, newLanguage)
	}
	return test
}

func FindLanguage(languageCode string) (language *Language) {
	var returnLanguage Language

	for _, singleLanguage := range languages {
		if singleLanguage.Code == languageCode {
			returnLanguage = singleLanguage
		}
	}
	return &returnLanguage
}

func FindAllLanguages() (languages []Language) {
	return languages
}

func UpdateLanguage(updatedLanguage Language) (reussi bool) {
	for i, singleLanguage := range languages {
		if singleLanguage.Code == updatedLanguage.Code {
			singleLanguage.Name = updatedLanguage.Name
			languages = append(languages[:i], singleLanguage)
			return true
		}
	}
	return false
}

func DeleteLanguage(languageCode string) (reussi bool) {
	for i, singleLanguage := range languages {
		if singleLanguage.Code == languageCode {
			languages = append(languages[:i], languages[i+1:]...)
			return true
		}
	}
	return false
}
