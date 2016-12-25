// Package errhandling defines error handling patterns. General rule, you should
// only define and export errors if users can act on these errors.
//
// Open Source Examples:
//	- None
//
// Articles and Videos:
//	- https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
//
package errhandling

import (
	"fmt"

	// This is a powerful library that allows you to add context to you errors.
	"github.com/pkg/errors"
)

/******************************************************************************
 * Sentinal Errors
 *****************************************************************************/
var ErrNotFound = errors.New("not found")

func ThrowSentinalError() error {
	return ErrNotFound
}

func Usage1() {
	err := ThrowSentinalError()

	// You can compare with an if statement
	if err == ErrNotFound {
		fmt.Println(err)
		return
	}

	// Or you can use a switch statement
	switch err {
	case ErrNotFound:
		// handle error
	case nil:
		// do nothing
	default:
		// handle error
	}
}

/******************************************************************************
 * Error Types
 *****************************************************************************/
type ErrValidation struct {
	FieldName string
	Cause     string
}

func (e ErrValidation) Error() string {
	return fmt.Sprintf("Invalid %v: %v", e.FieldName, e.Cause)
}

func ThrowErrorType() error {
	return ErrValidation{FieldName: "name", Cause: "empty string"}
}

func Usage2() {

}

/******************************************************************************
 * Inline Errors
 *****************************************************************************/
// This is the general case
// 	- if you expect users don't care about handling these specific errors.
func ThrowInlineError() error {
	return errors.New("something happend")
}

/******************************************************************************
 * Wrapping Errors
 *****************************************************************************/
// Dave Chaney's github.com/pkg/errors package allows for application to add
// context as errors are passed up the call stack.
//
// It enables K&D style errors: FAILED: WHY?: BECAUSE?
//
//	"could not read config: open failed: /tmp/.settings.xml: no such file"
//

func Usage3() {
	err := ThrowInlineError()
	err = errors.Wrap(err, "add context")

	// prints
	// "add context: something happend"
	fmt.Println(err)
}

/******************************************************************************
 * Error Tags
 *****************************************************************************/
// Error tags are more of a conventions. This helps developers track certain
// types of errors in their logs.

func Usage3() {
	// PROGERR: programmer errors, these errors should never occur if the
	// unless assumptions the develoopers made were incorrect.
	err := errors.New("progerr: invalid struct")

	// EDGECASE: edgecaes are expected but should not occur frequently
	// tagging this helps developers track the frequency of these errors
	//
	//	i.e. race conditions
	err = errors.New("edgecase: user not found")

	fmt.Println(err)
}
