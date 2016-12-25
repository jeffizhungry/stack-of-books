// Package apiclient define a pattern for creating client wrappers to external
// services that do not provide their own client libraries.
package apiclient

import "github.com/juju/ratelimit"

/******************************************************************************
 * Interfaces
 *****************************************************************************/

// You should ALWAYS define an interface for users of your package. This allows
// them use inject mocks for their unittests.
type Clienter interface {
	GetUsers() []string
	GetUserByID(id string) string
}

/******************************************************************************
 * Exported vs Unexported Types
 *****************************************************************************/

// Use an unexported type + interface if users never need to access or modify
// fields in the struct directly.
//
// This pattern is seen in https://github.com/Shopify/sarama
//
// Downside to this is that GoDocs is not as clean because you don't have
// hyperlinks to the implementation, since uneported structs are not shown in
// GoDocs.
type client struct {
	// You should store account specific data here
	token string

	// You should always have a rate limiter for each client. Usually external
	// services on a per account basis. This ratelimit bucket should be shared
	// between all API calls
	throttle *ratelimit.Bucket
}

func (c *client) GetUsers() []string {
	return []string{"Chad", "Amy"}
}

func (c *client) GetUserByID(id string) string {
	return "Randy"
}

// Use exported type if clients need to access exported fields
type ClientExported struct {
	AccessMe string
}

// Or if have multiple clients to choose from.
type ClientExportedToo struct {
	AccessMe string
}

/******************************************************************************
 * Request
 *****************************************************************************/

type Request struct {
}

/******************************************************************************
 * Constructors
 *****************************************************************************/

// You can return the interface instead of the concrete type.
//
// This allows you the flexibily to change implementation details without
// breaking users as long as the interface methods stay the same.
//
// This also forces users to use the interface, which they should if they want
// to mock out this client.
func NewClient(token string) Clienter {
	return &client{token: token}
}

// You can technically return unexported struct as well!
//
// This forces users to use this constructor method and prevents concreted
// type from being used (as fields to a struct, as args, as a return value)
//
// The main difference between this and returning the inteface is that it allow
// users to access exported fields. Honestly havent found a use case for
// this yet.
func NewClientUnexported(token string) *client {
	return &client{token: token}
}

/******************************************************************************
 * Mocks
 *****************************************************************************/

// Always provide mocks for your clients
type MockClient struct{}
