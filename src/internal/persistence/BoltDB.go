package persistence

import (
	"log"

	"github.com/boltdb/bolt"
)

// Open the my.db data file in your current directory.
var db, err = bolt.Open("my.db", 0600, nil)

func InitialiseDB() {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("language"))
		_, err = tx.CreateBucketIfNotExists([]byte("student"))

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}

func GetDataBase() *bolt.DB {
	return db
}

func CloseDatabase() {
	db.Close()
}
