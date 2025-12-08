// Package authapi maintains the web based api for auth access.
package authapi

import (
	"github.com/adrian-qorbani/atlas-service/api/http/api/middleware"
	"github.com/adrian-qorbani/atlas-service/app/api/auth"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Auth *auth.Auth
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {

	bearer := middleware.Bearer(cfg.Auth)
	basic := middleware.Basic(cfg.Auth)

	api := newAPI(cfg.Auth)

	app.HandleFunc("GET /auth/token/{kid}", api.token, basic)
	app.HandleFunc("GET /auth/authenticate", api.authenticate, bearer)
	app.HandleFunc("POST /auth/authorize", api.authorize)

}
