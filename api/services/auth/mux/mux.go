// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/adrian-qorbani/atlas-service/api/services/api/middleware"
	"github.com/adrian-qorbani/atlas-service/api/services/auth/routes/authapi"
	"github.com/adrian-qorbani/atlas-service/api/services/auth/routes/checkapi"
	"github.com/adrian-qorbani/atlas-service/business/api/auth"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
	"github.com/jmoiron/sqlx"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(build string, log *logger.Logger, db *sqlx.DB, auth *auth.Auth, shutdown chan os.Signal) *web.App {
	app := web.NewApp(shutdown, middleware.Logger(log), middleware.Errors(log), middleware.Metrics(), middleware.Panics())
	checkapi.Routes(build, app, log, db)
	authapi.Routes(app, auth)
	return app
}
