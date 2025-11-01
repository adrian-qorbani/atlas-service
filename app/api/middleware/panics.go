package middleware

import (
	"context"
	"fmt"
	"runtime/debug"
)

// Panics recover from panics and convert them to an error so it'll be reported in Metrics and
// handled in Errors.
func Panics(ctx context.Context, handler Handler) (err error) {

	// Defer a function to recover from a panic and set the err return
	// variable after the fact.
	defer func() {
		if rec := recover(); rec != nil {
			trace := debug.Stack()
			err = fmt.Errorf("PANIC [%v] TRACE [%s]", rec, string(trace))

			// to-do add metrics
		}
	}()

	return handler(ctx)
}
