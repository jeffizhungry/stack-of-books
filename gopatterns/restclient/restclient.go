// Package restclient define a pattern for creating client wrappers to external
// services that do not provide their own client libraries.
//
// Open Source Examples:
//	- https://github.com/kubernetes/kubernetes/tree/master/pkg/client/restclient
//	- https://github.com/stripe/stripe-go
//	- https://github.com/aws/aws-sdk-go
//
// Articles and Videos:
//	- Don't define domain structs yourself - https://mholt.github.io/json-to-go/
//
package restclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/juju/ratelimit"
)

/******************************************************************************
 * Interfaces
 *****************************************************************************/

// Define an interface for users, this provides a mocking point for them.
type Clienter interface {

	// This interface shouldn't expose any details that this is a rest
	// interface.
	GetUsers(ctx context.Context) []string
	GetUserByID(ctx context.Context, id string) User
}

/******************************************************************************
 * Client
 *****************************************************************************/

type client struct {
	// You should store account specific data here
	token    string
	username string

	// Instantiate your own http client, and set custom timeouts etc...
	httpClient *http.Client

	// You should always have a rate limiter for each client. Usually external
	// services on a per account basis. This ratelimit bucket should be shared
	// between all API calls.
	throttle *ratelimit.Bucket
}

// Context should be passed into your requests for logging, timeouts and
// cancellation signals.
func (c *client) GetUsers(ctx context.Context) []string {

	return []string{"Chad", "Amy"}
}

// See above
func (c *client) GetUserByID(ctx context.Context, id string) User {
	return User{Name: "Randy"}
}

// Common trick: This allows developers of this package to verify client
// satisfies the interface, before it is actually used.
var _ Clienter = (*client)(nil)

/******************************************************************************
 * Request
 *****************************************************************************/

// Clients make requests. Apps don't technically need to know about this.
type request struct {
}

func (r *request) Do(result interface{}) error {
	return nil
}

/******************************************************************************
 * Objects
 *****************************************************************************/
// Define stucts for payloads you expect to use and getback. These objects
// may be used by the rest of your application.
//
// Lookup an example for the REST service your are trying to connect to and
// use this to help you generate structs.
//
// https://mholt.github.io/json-to-go/
type User struct {

	// Keep members exported to users can read these directly without getter
	// and setter methods.
	Name    string    `json:"name"`
	Created time.Time `json:"created"`

	// Nested anonymous structs are sometimes easier to use
	Classes []struct {
		Period int    `json:"period"`
		Grade  string `json:"grade"`
	} `json:"classes"`

	// Define named structs for objects that may be decoupled from the
	// top level object. Depends on your app's usecase =D
	Friends []Friend `json:"friends"`
}

// It's useful to write custom methods on these object to help you application
// interpret the structure in certain ways.
func (u *User) IsFriendsWith(name string) bool {
	return true
}

type Friend struct {
	Name               int    `json:"name"`
	RelationshipStatus string `json:"relationship"`
}

/******************************************************************************
 * Your Domain Objects
 *****************************************************************************/
//
// DONT DO THIS
//
// You should not be mixing your domain objects with the objects of the service
// you are connecting to in this package.
type RESTObject struct {
	YourDomainObject struct{}
}

/******************************************************************************
 * Errors
 *****************************************************************************/
// It is useful to define errors if users of this package can act of them.
var (
	ErrNotFound = errors.New("object not found")

	// Can be signal to reauthenticate client
	ErrNotAutorized = errors.New("not authroized")

	// May be less useful. This can be retried in the client.
	ErrTooManyRequests = errors.New("too many request")
)

// If the REST API has an error response payload, you might want to use an
// error type instead.
type ErrResp struct {
	Code int
	Body string
}

func (e ErrResp) Error() string {
	return fmt.Sprintf("<blank> API error: code=%v, body=%v", e.Code, e.Body)
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

/******************************************************************************
 * Mocks
 *****************************************************************************/

// Always provide mocks for your clients
type MockClient struct{}
