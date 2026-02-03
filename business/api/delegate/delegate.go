// Package delegate provides the ability to make function calls between
// different domain packages when an import is not possible.
package delegate

import "github.com/adrian-qorbani/atlas-service/foundation/logger"

// These types are just for documentation so we know what keys go
// where in the map.
type (
	domain string
	action string
)

// Delegate manages the set of functions to be called by domain
// packages when an import is not possible.
type Delegate struct {
	log   *logger.Logger
	funcs map[domain]map[action][]Func
}
