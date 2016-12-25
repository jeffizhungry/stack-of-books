// Package messages
//
// Open Source Examples:
//	- None
//
// Articles and Videos:
//	- CorrelationIDs - Building Microservices by Sam Newman
//
package messages

import "time"

// Header should be added to your messages if your transport mechanism doesn't
// already provide these fields.
type Header struct {

	// MessageID should be alphanumeric string unique to this message.
	MessageID string

	// CorrelationID is an alphanumerica string that should be generated
	// once when the first message is created (typically at your api gateway)
	// and passed through as more events an messages are generated in response.
	CorrelationID string

	// Created helps timestamp and identify staleness of this message
	Created time.Time
}

// Define your message type
type Metric struct {
	Header
	Name  string
	Value float64
}

// Messages should be validated before they are sent / received.
func (m *Metric) Validate() error {

}

// Custom encode should be defined incase you want to switch out the
// serialization. It is not uncommon to start with JSON and switch to protobufs
// or avro later on for more efficieny.
func (m *Metric) Encode() ([]byte, error) {

}

// See encode
func (m *Metric) Decode(data []byte) error {

}

// Custom string functions help with debugging.
func (m *Metric) String() string {
	// return utils.Prettify(t)
}
