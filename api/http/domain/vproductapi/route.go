package vproductapi

import (
	"github.com/adrian-qorbani/atlas-service/api/http/api/middleware"
	"github.com/adrian-qorbani/atlas-service/app/api/auth"
	"github.com/adrian-qorbani/atlas-service/app/api/authclient"
	"github.com/adrian-qorbani/atlas-service/app/domain/vproductapp"
	"github.com/adrian-qorbani/atlas-service/business/domain/userbus"
	"github.com/adrian-qorbani/atlas-service/business/domain/vproductbus"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log         *logger.Logger
	UserBus     *userbus.Business
	VProductBus *vproductbus.Business
	AuthClient  *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := middleware.Authenticate(cfg.Log, cfg.AuthClient)
	ruleAdmin := middleware.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAdminOnly)

	api := newAPI(vproductapp.NewApp(cfg.VProductBus))
	app.HandleFunc("GET /vproducts", api.query, authen, ruleAdmin)
}
