// Package defines pattern for handling domain objects or models.
//
// Models should be decoupled with the database abstraction logic. Frameworks
// such as Django or Ruby on Rails where objects "save themselves", typically use
// global database connections and do not allow to easily allow mocking
// external services.
package database

import (
	"database/sql"
	"errors"
)

// Sometimes it's useful to define a not found error, because this is a common
// issue when working with any storage layer.
var ErrNotFound = errors.New("model not found")

// Define an interface so it can be a mocking point
type StudentDALer interface {
	GetStudentByName(name string) (*Student, error)
	GetStudents(name string) ([]*Student, error)
	PutStudent(s *Student) error
	UpdateStudentName(id string, name string) error
}

type StudentDAL struct {
	conn *sql.DB
}

func (s *StudentDAL) GetStudentByName(name string) (*Student, error) {
	return nil, nil
}

func (s *StudentDAL) GetStudents() ([]*Student, error) {
	return nil, nil
}

type Student struct{}

//
func (s *Student) Validate() error {

}
