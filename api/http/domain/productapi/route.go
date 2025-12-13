package productapi

import (
	"github.com/adrian-qorbani/atlas-service/api/http/api/middleware"
	"github.com/adrian-qorbani/atlas-service/app/api/auth"
	"github.com/adrian-qorbani/atlas-service/app/api/authclient"
	"github.com/adrian-qorbani/atlas-service/app/domain/productapp"
	"github.com/adrian-qorbani/atlas-service/business/domain/productbus"
	"github.com/adrian-qorbani/atlas-service/business/domain/userbus"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	UserBus    *userbus.Business
	ProductBus *productbus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := middleware.Authenticate(cfg.Log, cfg.AuthClient)
	ruleAny := middleware.Authorize(cfg.Log, cfg.AuthClient, auth.RuleAny)
	ruleUserOnly := middleware.Authorize(cfg.Log, cfg.AuthClient, auth.RuleUserOnly)
	ruleAuthorizeProduct := middleware.AuthorizeProduct(cfg.Log, cfg.AuthClient, cfg.ProductBus)

	api := newAPI(productapp.NewApp(cfg.ProductBus))
	app.HandleFunc("GET /products", api.query, authen, ruleAny)
	app.HandleFunc("GET /products/{product_id}", api.queryByID, authen, ruleAuthorizeProduct)
	app.HandleFunc("POST /products", api.create, authen, ruleUserOnly)
	app.HandleFunc("PUT /products/{product_id}", api.update, authen, ruleAuthorizeProduct)
	app.HandleFunc("DELETE /products/{product_id}", api.delete, authen, ruleAuthorizeProduct)
}
