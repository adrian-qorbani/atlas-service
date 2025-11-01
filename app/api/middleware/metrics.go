package middleware

import (
	"context"

	"github.com/adrian-qorbani/atlas-service/app/api/metrics"
)

// Metrics updates program counters
func Metrics(ctx context.Context, handler Handler) error {
	ctx = metrics.Set(ctx)

	err := handler(ctx)

	n := metrics.AddRequests(ctx)

	// to be updated; 1000 too small for production
	if n%1000 == 0 {
		metrics.AddGoroutines(ctx)
	}

	if err != nil {
		metrics.AddErrors(ctx)
	}

	return err
}
