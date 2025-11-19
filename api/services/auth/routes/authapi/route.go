// Package authapi maintains the web based api for auth access.
package authapi

import (
	"github.com/adrian-qorbani/atlas-service/api/services/api/middleware"
	"github.com/adrian-qorbani/atlas-service/business/api/auth"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App, a *auth.Auth) {
	authen := middleware.Authenticate(a)

	api := newAPI(a)
	app.HandleFunc("GET /auth/token/{kid}", api.token, authen)
	app.HandleFunc("GET /auth/authenticate", api.authenticate, authen)
	app.HandleFunc("POST /auth/authorize", api.authorize)

}
