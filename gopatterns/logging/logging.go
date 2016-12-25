// Package logging attempts to define logging practices using a structured logger
// such as logrus
//
// Open Source Examples:
//	- github.com/Sirupsen/logrus
//
// Articles and Videos:
//	- http://dev.splunk.com/view/logging-best-practices/SP-CAAADP6
//
package logging

import (
	"errors"

	"github.com/Sirupsen/logrus"
)

const contextName = "lib.logging"

func Usage() {

	// Always add request specific information to your context logger
	user := "Jeff"
	log := logrus.WithFields(logrus.Fields{
		"context": contextName,
		"user":    user,
	})

	// DEBUG: Useful for dev or high volume debugging, should be turned off in prod.
	// Otherwise this will flood your logs.
	log.Debug("received user request")

	// INFO: Notice something happend
	log.Info("application started")

	// WARN: Error occured, but we can continue
	log.Warn("something happend")

	// ERROR: Error occured, and we shouldn't continue
	log.Error("uh oh...")

	// Use WithError to track any errors
	err := errors.New("issue")
	log.WithError(err).Error("describe what threw the error")
}

/******************************************************************************
 * Intialized log
 *****************************************************************************/

type Consumer struct {
	log *logrus.Entry
}

func NewConsumer() *Consumer {
	return &Consumer{}
}
