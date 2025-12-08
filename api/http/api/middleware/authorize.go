package middleware

import (
	"context"
	"net/http"

	"github.com/adrian-qorbani/atlas-service/app/api/authclient"
	"github.com/adrian-qorbani/atlas-service/app/api/middleware"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Authorize executes the authorize middleware functionality
func Authorize(log *logger.Logger, client *authclient.Client, rule string) web.MidFunc {

	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				return handler(ctx, w, r)
			}
			return middleware.Authorize(ctx, log, client, rule, hdl)
		}
		return h
	}
	return m
}
