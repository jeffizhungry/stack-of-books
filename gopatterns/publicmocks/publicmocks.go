// Package publicmocks encourages the philosophy that mocks should be provided
// as part of the public inteface to a library.
//
// 	Video: 		http://www.youtube.com/watch?v=yszygk1cpEc&t=54m14s
// 	Seen here: 	https://github.com/Shopify/sarama
//
// Problem
//
// Time and time again, developers have experienced that they need to create
// mocks for libraries they use. Since this is repeated code why not have it
// as part of the library.
//
// One concerned is that this forces users to use a specific mocking
// framework however, these mocks are optional. If you don't like the mocks
// defined this library, then just make your own mock!
package publicmocks

/******************************************************************************
 * Types
 *****************************************************************************/

// Define the interface, because interfaces are mocking points in the code.
type Studenter interface {
	DoHomework()
	GetHomework() Homework
}

type Homework struct{}

// Define the actual type.
type Student struct{}

func (s *Student) DoHomework() {
}

func (s *Student) GetHomework() Homework {
	return Homework{}
}

/******************************************************************************
 * Mockery for Mock Generation
 *****************************************************************************/

// Mockery is a mock generation utility for http://github.com/testify/mocks
//
// The project can be found here: https://github.com/vektra/mockery
//
// Installation:
//	- go get github.com/vektra/mockery/.../
//
// You can call have go generate call mockery to generate your code for you
// To use go generate you literally put the comment below in your code.

// go:generate mockery -name Studenter

// This will put your mocks in the yourlib/mocks directory
