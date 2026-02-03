package delegate

import "context"

// Func represents a function that is registered and called by the system.
type Func func(context.Context, Data) error

// Data represents an event between domains.
type Data struct {
	Domain    string
	Action    string
	RawParams []byte
}
