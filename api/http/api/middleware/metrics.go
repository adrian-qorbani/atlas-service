package middleware

import (
	"context"
	"net/http"

	"github.com/adrian-qorbani/atlas-service/app/api/middleware"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Metrics update program counters using middleware functionality
func Metrics() web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}
			return middleware.Metrics(ctx, hdl)
		}
		return h
	}
	return m
}
