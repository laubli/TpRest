package persistence

import (
	. "internal/entities/language"
	"sort"
)

type LanguageDAOMemory struct{}

//var _ Languages = (*LanguageDAOMemory)(nil)

//var mapLanguage = make(map[string]Language)

var (
	languages []Language // slice (dynamically sized array)
)

func init() {
	languages = []Language{Language{
		Code: "1",
		Name: "Nom 1",
	}}
}


func (l LanguageDAOMemory) FindAll() []Language {
	sort.Sort(LanguageList(languages))
	return languages
}

func (l LanguageDAOMemory) Find(languageCode string) *Language {
	var returnLanguage Language

	for _, singleLanguage := range languages {
		if singleLanguage.Code == languageCode {
			returnLanguage = singleLanguage
		}
	}
	return &returnLanguage
}

func (l LanguageDAOMemory) Exist(testLanguage string) bool {
	for _, singleLanguage := range languages {
		if singleLanguage.Code == testLanguage {
			return true
		}
	}
	return false
}

func (l LanguageDAOMemory) Delete(languageCode string) bool {
	for i, singleLanguage := range languages {
		if singleLanguage.Code == languageCode {
			languages = append(languages[:i], languages[i+1:]...)
			return true
		}
	}
	return false
}

func (l LanguageDAOMemory) Create(newLanguage Language) bool {
	if !l.Exist(newLanguage.Code) {
		languages = append(languages, newLanguage)
		return true
	}
	return false
}

func (l LanguageDAOMemory) Update(updatedLanguage Language) bool {
	for i, singleLanguage := range languages {
		if singleLanguage.Code == updatedLanguage.Code {
			singleLanguage.Name = updatedLanguage.Name
			languages = append(languages[:i], singleLanguage)
			return true
		}
	}
	return false
}

func (l LanguageDAOMemory) NewLanguageDAOMemory() LanguageDAOMemory {
	return LanguageDAOMemory{}
}
