package mux

import (
	"context"

	"github.com/adrian-qorbani/atlas-service/api/http/api/domain/authapi"
	"github.com/adrian-qorbani/atlas-service/api/http/api/domain/checkapi"
	"github.com/adrian-qorbani/atlas-service/api/http/api/middleware"
	"github.com/adrian-qorbani/atlas-service/app/api/auth"
	"github.com/adrian-qorbani/atlas-service/app/api/authclient"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
	"github.com/jmoiron/sqlx"
)

// WebAPISales constructs a http.Handler with all application routes bound.
func WebAPISales(build string, log *logger.Logger, db *sqlx.DB, authClient *authclient.Client) *web.App {

	logger := func(ctx context.Context, msg string, v ...any) {
		log.Info(ctx, msg, v...)
	}

	app := web.NewApp(logger, middleware.Logger(log), middleware.Errors(log), middleware.Metrics(), middleware.Panics())
	checkapi.Routes(build, app, log, db)

	return app
}

// WebAPIAuth constructs a http.Handler with all application routes bound.
func WebAPIAuth(build string, log *logger.Logger, db *sqlx.DB, auth *auth.Auth) *web.App {

	logger := func(ctx context.Context, msg string, v ...any) {
		log.Info(ctx, msg, v...)
	}

	app := web.NewApp(logger, middleware.Logger(log), middleware.Errors(log), middleware.Metrics(), middleware.Panics())
	checkapi.Routes(build, app, log, db)
	authapi.Routes(app, auth)
	return app
}
