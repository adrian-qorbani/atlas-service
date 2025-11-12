package middleware

import (
	"context"
	"net/http"

	"github.com/adrian-qorbani/atlas-service/app/api/middleware"
	"github.com/adrian-qorbani/atlas-service/business/api/auth"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Authorize executes the authorize middleware functionality
func Authorize(auth *auth.Auth, rule string) web.MidFunc {

	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}
			return middleware.Authorize(ctx, auth, rule, hdl)
		}
		return h
	}
	return m
}
