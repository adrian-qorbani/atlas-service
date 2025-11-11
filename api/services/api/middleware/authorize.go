package middleware

import (
	"context"
	"net/http"

	"github.com/adrian-qorbani/atlas-service/app/api/middleware"
	"github.com/adrian-qorbani/atlas-service/business/api/auth"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Authorize executes the authorize middleware functionality
func Authorize(auth *auth.Auth) web.MidFunc {

	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}
			return middleware.Authorization(ctx, auth, r.Header.Get("authorization"), hdl)
		}
		return h
	}
	return m
}
