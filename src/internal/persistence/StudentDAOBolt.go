package persistence

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	. "internal/entities/student"
	"sort"

	"github.com/boltdb/bolt"
)

type StudentDAOBolt struct{}

var studentBucketName = "student"

var _ StudentDAOInterface = (*StudentDAOBolt)(nil)

func (l StudentDAOBolt) FindAll() []Student {
	var Students = []Student{}

	GetDataBase().View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(studentBucketName))

		b.ForEach(func(k, v []byte) error {
			var tmpStudent = Student{}
			err = json.Unmarshal(v, &tmpStudent)
			Students = append(Students, tmpStudent)
			return nil
		})

		if err != nil {
			fmt.Printf("Error: cannot unmarshall Student %s", err)
		}

		return nil
	})

	sort.Sort(StudentList(Students))
	return Students
}

func (l StudentDAOBolt) Find(id int) *Student {
	var Student = Student{}

	if !l.Exist(id) {
		return nil
	}

	GetDataBase().View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(studentBucketName))

		var byteStudent = b.Get([]byte(string(id)))
		err = json.Unmarshal(byteStudent, &Student)

		if err != nil {
			fmt.Printf("Error: cannot unmarshall Student %s", err)
		}

		return nil
	})

	return &Student
}

func (l StudentDAOBolt) Exist(id int) bool {
	return GetDataBase().View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(studentBucketName))

		// Search if Student already exist, if not send error
		var alreadyExistingKey, _ = b.Cursor().Seek([]byte(string(id)))
		if bytes.Compare(alreadyExistingKey, []byte(string(id))) != 0 {
			return errors.New("Cannot find matching key in bucket " + studentBucketName) // FIXME non-exploited error
		}

		return nil
	}) == nil
}

func (l StudentDAOBolt) Delete(id int) bool {
	return GetDataBase().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(studentBucketName))

		// Search if Student already exist, if not send error
		if !l.Exist(id) {
			return errors.New("The Student you want to delete doesn't exist in bucket " + studentBucketName)
		}

		return b.Delete([]byte(string(id)))
	}) == nil
}

func (l StudentDAOBolt) Create(Student Student) bool {
	return GetDataBase().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(studentBucketName))

		buf, err := json.Marshal(Student)
		if err != nil {
			return err
		}

		// Search if Student already exist, if so send error
		if l.Exist(Student.Id) {
			return errors.New("Student with the same key already present in bucket " + studentBucketName)
		}

		return b.Put([]byte(string(Student.Id)), buf)
	}) == nil
}

func (l StudentDAOBolt) Update(Student Student) bool {
	return GetDataBase().Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(studentBucketName))

		buf, err := json.Marshal(Student)
		if err != nil {
			return err
		}

		// Search if Student already exist, if not send error
		if !l.Exist(Student.Id) {
			return errors.New("The Student you want to update doesn't exist in bucket " + studentBucketName)
		}

		return b.Put([]byte(string(Student.Id)), buf)
	}) == nil
}

func (l StudentDAOBolt) NewStudentDAOMemory() StudentDAOMemory { // FIXME: Should NOT be there, modify interface
	return StudentDAOMemory{}
}
