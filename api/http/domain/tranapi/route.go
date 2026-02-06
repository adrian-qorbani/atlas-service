package tranapi

import (
	"github.com/adrian-qorbani/atlas-service/api/http/api/middleware"
	"github.com/adrian-qorbani/atlas-service/app/api/authclient"
	tranapp "github.com/adrian-qorbani/atlas-service/app/domain/transapp"
	"github.com/adrian-qorbani/atlas-service/business/api/sqldb"
	"github.com/adrian-qorbani/atlas-service/business/domain/productbus"
	"github.com/adrian-qorbani/atlas-service/business/domain/userbus"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Log        *logger.Logger
	DB         *sqlx.DB
	UserBus    *userbus.Business
	ProductBus *productbus.Business
	AuthClient *authclient.Client
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := middleware.Authenticate(cfg.Log, cfg.AuthClient)
	transaction := middleware.BeginCommitRollback(cfg.Log, sqldb.NewBeginner(cfg.DB))

	api := newAPI(tranapp.NewApp(cfg.UserBus, cfg.ProductBus))
	app.HandleFunc("POST /tranexample", api.create, authen, transaction)
}
