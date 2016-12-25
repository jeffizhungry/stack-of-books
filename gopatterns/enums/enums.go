// Package enums attempts to should show conventions of defining enums in Go,
// because the language has terrible support for enums.
//
// Open Source Examples:
//	- None
//
// Articles and Videos:
//	- None
//
package enums

/******************************************************************************
 * Enums that DON'T exit the scope of this codebase
 *****************************************************************************/
// Enums that are strictly used in this scope of this codebase should be defined
// as an int type.

type ClientType int

// Enums should generally be prefixed with the enum type:
//
// - If there are multiple enums defined in the same package,
// - or if you expect to define more enum types in the same package
const (
	ClientTypeUser ClientType = iota
	ClientTypeAdmin
)

/******************************************************************************
 * Enums that DO exit the scope of this codebase
 *****************************************************************************/
// Enums that are going to be sent to another service, i.e. returned in an
// REST handler, or stored in a database should be defined as an string type.
//
// If this is the usecase, you are tyrpically going to need to convert it to
// a string type anyways. So rather then defining a method to make that
// conversion, just make the type a string!
type ClientStatus string

const (
	ClientStatusOK       ClientStatus = "OK"
	ClientStatusUpdating ClientStatus = "Updating"
	ClientStatusInvalid  ClientStatus = "Invalid"
)
