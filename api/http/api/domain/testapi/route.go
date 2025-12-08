package testapi

import (
	"github.com/adrian-qorbani/atlas-service/api/http/api/middleware"
	"github.com/adrian-qorbani/atlas-service/app/api/auth"
	"github.com/adrian-qorbani/atlas-service/app/api/authclient"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Config contains all mandatory systems required by handlers
type Config struct {
	Log        *logger.Logger
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group
func Routes(app *web.App, cfg Config) {

	authen := middleware.Authenticate(cfg.Log, cfg.AuthClient)
	athAdminOnly := middleware.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAdminOnly)

	api := newAPI()

	app.HandleFunc("GET /testerror", api.testError)
	app.HandleFunc("GET /testpanic", api.testPanic)
	app.HandleFunc("GET /testauth", api.testGood, authen, athAdminOnly)

}
