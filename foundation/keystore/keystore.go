// Package keystore implements the auth.KeyLookup interface. This implements
// an in-memory keystore for JWT support.
package keystore

import (
	"errors"
	"sync"
)

// ErrKeyNotFound is returned when a key identified by a kid is not found.
var ErrKeyNotFound = errors.New("key not found")

// key represents key information.
type key struct {
	privatePEM string
	publicPEM  string
}

// KeyStore represents an in memory store implementation of the
// KeyLookup interface for use with the auth package.
type KeyStore struct {
	store map[string]key
	mu    sync.RWMutex
}
