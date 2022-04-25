package persistence

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	. "internal/entities/language"
	"sort"

	"github.com/boltdb/bolt"
)

type LanguageDAOBolt struct{}

var languageBucketName = "language"

var _ Languages = (*LanguageDAOBolt)(nil)

func (l LanguageDAOBolt) FindAll() []Language {
	var languages = []Language{}

	GetDataBase().View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(languageBucketName))

		b.ForEach(func(k, v []byte) error {
			var tmpLanguage = Language{}
			err = json.Unmarshal(v, &tmpLanguage)
			languages = append(languages, tmpLanguage)
			return nil
		})

		if err != nil {
			fmt.Printf("Error: cannot unmarshall language %s", err)
		}

		return nil
	})

	sort.Sort(LanguageList(languages))
	return languages
}

func (l LanguageDAOBolt) Find(code string) *Language {
	var language = Language{}

	if !l.Exist(code) {
		return nil
	}

	GetDataBase().View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(languageBucketName))

		var byteLanguage = b.Get([]byte(code))
		err = json.Unmarshal(byteLanguage, &language)

		if err != nil {
			fmt.Printf("Error: cannot unmarshall language %s", err)
		}

		return nil
	})

	return &language
}

func (l LanguageDAOBolt) Exist(code string) bool {
	return GetDataBase().View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(languageBucketName))

		// Search if language already exist, if not send error
		var alreadyExistingKey, _ = b.Cursor().Seek([]byte(code))
		if bytes.Compare(alreadyExistingKey, []byte(code)) != 0 {
			return errors.New("Cannot find matching key in bucket " + languageBucketName) // FIXME non-exploited error
		}

		return nil
	}) == nil
}

func (l LanguageDAOBolt) Delete(code string) bool {
	return GetDataBase().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(languageBucketName))

		// Search if language already exist, if not send error
		if !l.Exist(code) {
			return errors.New("The language you want to delete doesn't exist in bucket " + languageBucketName)
		}

		return b.Delete([]byte(code))
	}) == nil
}

func (l LanguageDAOBolt) Create(Language Language) bool {
	return GetDataBase().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(languageBucketName))

		buf, err := json.Marshal(Language)
		if err != nil {
			return err
		}

		// Search if language already exist, if so send error
		if l.Exist(Language.Code) {
			return errors.New("Language with the same key already present in bucket " + languageBucketName)
		}

		return b.Put([]byte(Language.Code), buf)
	}) == nil
}

func (l LanguageDAOBolt) Update(Language Language) bool {
	return GetDataBase().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(languageBucketName))

		buf, err := json.Marshal(Language)
		if err != nil {
			return err
		}

		// Search if language already exist, if not send error
		if !l.Exist(Language.Code) {
			return errors.New("The language you want to update doesn't exist in bucket " + languageBucketName)
		}

		return b.Put([]byte(Language.Code), buf)
	}) == nil
}

func (l LanguageDAOBolt) NewLanguageDAOMemory() LanguageDAOMemory { // FIXME: Should NOT be there, modify interface
	return LanguageDAOMemory{}
}
