package homeapi

import (
	"github.com/adrian-qorbani/atlas-service/api/http/api/middleware"
	"github.com/adrian-qorbani/atlas-service/app/api/auth"
	"github.com/adrian-qorbani/atlas-service/app/api/authclient"
	"github.com/adrian-qorbani/atlas-service/app/domain/homeapp"
	"github.com/adrian-qorbani/atlas-service/business/domain/homebus"
	"github.com/adrian-qorbani/atlas-service/business/domain/userbus"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	UserBus    *userbus.Business
	HomeBus    *homebus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {

	authen := middleware.Authenticate(cfg.Log, cfg.AuthClient)
	ruleAny := middleware.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAny)
	ruleUserOnly := middleware.Authorize(cfg.Log, cfg.AuthClient, auth.RuleUserOnly)
	ruleAuthorizeHome := middleware.AuthorizeHome(cfg.Log, cfg.AuthClient, cfg.HomeBus)

	api := newAPI(homeapp.NewApp(cfg.HomeBus))
	app.HandleFunc("GET /homes", api.query, authen, ruleAny)
	app.HandleFunc("GET /homes/{home_id}", api.queryByID, authen, ruleAuthorizeHome)
	app.HandleFunc("POST /homes", api.create, authen, ruleUserOnly)
	app.HandleFunc("PUT /homes/{home_id}", api.update, authen, ruleAuthorizeHome)
	app.HandleFunc("DELETE /homes/{home_id}", api.delete, authen, ruleAuthorizeHome)

	// 	app.HandleFunc("GET /users", api.query, authen, ruleAdmin)
	// app.HandleFunc("GET /users/{user_id}", api.queryByID, authen, ruleAuthorizeUser)
}
