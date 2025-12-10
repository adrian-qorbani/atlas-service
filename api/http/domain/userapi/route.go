package userapi

import (
	"github.com/adrian-qorbani/atlas-service/api/http/api/middleware"
	"github.com/adrian-qorbani/atlas-service/app/api/auth"
	"github.com/adrian-qorbani/atlas-service/app/api/authclient"
	"github.com/adrian-qorbani/atlas-service/app/domain/userapp"
	"github.com/adrian-qorbani/atlas-service/business/domain/userbus"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	UserBus    *userbus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {

	authen := middleware.Authenticate(cfg.Log, cfg.AuthClient)
	ruleAdmin := middleware.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAdminOnly)
	ruleAuthorizeUser := middleware.AuthorizeUser(cfg.Log, cfg.AuthClient, cfg.UserBus, auth.RuleAdminOrSubject)
	ruleAuthorizeAdmin := middleware.AuthorizeUser(cfg.Log, cfg.AuthClient, cfg.UserBus, auth.RuleAdminOnly)

	api := newAPI(userapp.NewApp(cfg.UserBus))
	app.HandleFunc("GET /users", api.query, authen, ruleAdmin)
	app.HandleFunc("GET /users/{user_id}", api.queryByID, authen, ruleAuthorizeUser)
	app.HandleFunc("POST /users", api.create, authen, ruleAdmin)
	app.HandleFunc("PUT /users/role/{user_id}", api.updateRole, authen, ruleAuthorizeAdmin)
	app.HandleFunc("PUT /users/{user_id}", api.update, authen, ruleAuthorizeUser)
	app.HandleFunc("DELETE /users/{user_id}", api.delete, authen, ruleAuthorizeUser)
}
