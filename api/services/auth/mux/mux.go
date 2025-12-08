// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"context"

	"github.com/adrian-qorbani/atlas-service/api/services/api/middleware"
	"github.com/adrian-qorbani/atlas-service/api/services/auth/routes/authapi"
	"github.com/adrian-qorbani/atlas-service/api/services/auth/routes/checkapi"
	"github.com/adrian-qorbani/atlas-service/app/api/auth"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
	"github.com/jmoiron/sqlx"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(build string, log *logger.Logger, db *sqlx.DB, auth *auth.Auth) *web.App {

	logger := func(ctx context.Context, msg string, v ...any) {
		log.Info(ctx, msg, v...)
	}

	app := web.NewApp(logger, middleware.Logger(log), middleware.Errors(log), middleware.Metrics(), middleware.Panics())
	checkapi.Routes(build, app, log, db)
	authapi.Routes(app, auth)
	return app
}
