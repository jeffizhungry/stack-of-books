// Package database shows a pattern for objects interacting with persistent storage.
//
// Models should be decoupled with the database abstraction layer logic. Frameworks
// such as Django or Ruby on Rails where objects "save themselves", typically use
// global database connections and do not allow to easily allow mocking external
// services.
//
// Statically typed languages need to use dependency injection to allow for mocking.
//
// Open Source Examples:
//	- None
//
// Articles and Videos:
//	- None
//
package database

import (
	"database/sql"
	"errors"
)

// Sometimes it's useful to define a not found error, because this is a common
// issue when working with any storage layer.
var ErrNotFound = errors.New("model not found")

/******************************************************************************
 * Abstract type
 *****************************************************************************/

// Define an interface so it can be a mocking point
type StudentDALer interface {
	// Try to be consistent with you CRUD naming conventions. It doesn't matter
	// what it is but just be consistent!
	//
	//	Typically operations you might need are:
	//
	//		Get<model>			- retrieve object by primary key
	//		Get<model>ByName	- retrieve object by secondary index
	//		Put<model>			- insert or update a full model
	//		Update<model>       - update select fields in the model
	//		Delete<model>		- delete model by primary key
	//		List<model>			- list models that need primary index
	//		List<model>ByName 	- list models by some other criteria
	//
	//	Notes: This interface should show NO KNOWLEDGE of the underlying storage
	//	Arguments and return values must always be:
	//
	//		Primitive types		- i.e. string, int, error...
	//		User Defined type   - i.e. Student
	//
	GetStudentByName(name string) (*Student, error)
	GetStudents() ([]*Student, error)
	PutStudent(s *Student) error
	UpdateStudentName(id, name string) error
}

/******************************************************************************
 * Concrete type interacting with real persistent storage
 *****************************************************************************/

// Define the concrete type that actually interacts with your storage layer.
type StudentDAL struct {
	conn *sql.DB
}

var cache interface {
	Get(name string) (*Student, error)
}

func (s *StudentDAL) GetStudentByName(name string) (*Student, error) {

	// If you have a caching layer, this would be good place to put this logic
	//
	// It's difficult to use a caching decorator like you would in Python or
	// Java, because Go does not have a concept of passing around arbitrary
	// arguements without losing static typicing.
	res, err := cache.Get(name)
	if err == nil && res != nil {
		return res, err
	}

	// Get your object
	//
	// res, err := conn.Query()

	// Create or unmarshal your object
	//
	// var student Student
	// if err := json.Unmarshal(res, &student); err != nil {
	//	 return nil, err
	// }

	// You should only return valid objects in your business logic
	//
	// Technically you shouldn't need to do this. This is a programmer error.
	// Because if you called validate before saving this object, and there are
	// no issues with you database library, then this should never happen.
	if err := s.Validate(); err != nil {
		return err
	}
	return nil, nil
}

func (s *StudentDAL) GetStudents() ([]*Student, error) {
	return nil, nil
}

func (s *StudentDAL) PutStudent(s *Student) error {
	// You should only put valid objects in your database
	if err := s.Validate(); err != nil {
		return err
	}
	return nil
}

func (s *StudentDAL) UpdateStudentName(id, name string) error {
	return nil
}

// This is a common pattern. Use compiler to verify your concrete type
// satisfies the interface you defined. This will help detect errors before
// developers try to use this package.
var _ StudentDALer = (*StudentDAL)(nil)

/******************************************************************************
 * Models
 *****************************************************************************/

type Student struct{}

// Models should always define a validate function to validate before inserts.
func (s *Student) Validate() error {

}
