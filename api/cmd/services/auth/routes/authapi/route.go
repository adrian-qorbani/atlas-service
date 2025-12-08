// Package authapi maintains the web based api for auth access.
package authapi

import (
	"github.com/adrian-qorbani/atlas-service/api/services/api/middleware"
	"github.com/adrian-qorbani/atlas-service/app/api/auth"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App, ath *auth.Auth) {

	bearer := middleware.Bearer(ath)
	basic := middleware.Basic(ath)

	api := newAPI(ath)

	app.HandleFunc("GET /auth/token/{kid}", api.token, basic)
	app.HandleFunc("GET /auth/authenticate", api.authenticate, bearer)
	app.HandleFunc("POST /auth/authorize", api.authorize)

}
