package persistence

import . "internal/entities/language"

type Languages interface {
	FindAll() []Language
	Find(code string) *Language
	Exist(code string) bool
	Delete(code string) bool
	Create(Language Language) bool
	Update(Language Language) bool
	NewLanguageDAOMemory() LanguageDAOMemory // TODO should be renamed
}
