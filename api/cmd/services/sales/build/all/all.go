// Package all binds all the routes into the specified app.
package all

import (
	"github.com/adrian-qorbani/atlas-service/api/http/api/domain/checkapi"
	"github.com/adrian-qorbani/atlas-service/api/http/api/domain/testapi"
	"github.com/adrian-qorbani/atlas-service/api/http/api/mux"
	"github.com/adrian-qorbani/atlas-service/api/http/domain/homeapi"
	"github.com/adrian-qorbani/atlas-service/api/http/domain/productapi"
	"github.com/adrian-qorbani/atlas-service/api/http/domain/userapi"
	"github.com/adrian-qorbani/atlas-service/api/http/domain/vproductapi"
	"github.com/adrian-qorbani/atlas-service/business/api/delegate"
	"github.com/adrian-qorbani/atlas-service/business/domain/homebus"
	"github.com/adrian-qorbani/atlas-service/business/domain/homebus/stores/homedb"
	"github.com/adrian-qorbani/atlas-service/business/domain/productbus"
	"github.com/adrian-qorbani/atlas-service/business/domain/productbus/stores/productdb"
	"github.com/adrian-qorbani/atlas-service/business/domain/userbus"
	"github.com/adrian-qorbani/atlas-service/business/domain/userbus/store/userdb"
	"github.com/adrian-qorbani/atlas-service/business/domain/vproductbus"
	"github.com/adrian-qorbani/atlas-service/business/domain/vproductbus/stores/vproductdb"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// Routes constructs the add value which provides the implementation of
// of RouteAdder for specifying what routes to bind to this instance.
func Routes() add {
	return add{}
}

type add struct{}

// Add implements the RouterAdder interface.
func (add) Add(app *web.App, cfg mux.Config) {

	// Construct the business domain packages we need here so we are using the
	// sames instances for the different set of domain apis.
	delegate := delegate.New(cfg.Log)
	userBus := userbus.NewBusiness(cfg.Log, delegate, userdb.NewStore(cfg.Log, cfg.DB))
	homeBus := homebus.NewBusiness(cfg.Log, userBus, delegate, homedb.NewStore(cfg.Log, cfg.DB))
	productBus := productbus.NewBusiness(cfg.Log, userBus, delegate, productdb.NewStore(cfg.Log, cfg.DB))
	vproductBus := vproductbus.NewBusiness(vproductdb.NewStore(cfg.Log, cfg.DB))

	checkapi.Routes(app, checkapi.Config{
		Build: cfg.Build,
		Log:   cfg.Log,
		DB:    cfg.DB,
	})

	testapi.Routes(app, testapi.Config{
		Log:        cfg.Log,
		AuthClient: cfg.AuthClient,
	})

	userapi.Routes(app, userapi.Config{
		Log:        cfg.Log,
		UserBus:    userBus,
		AuthClient: cfg.AuthClient,
	})

	homeapi.Routes(app, homeapi.Config{
		Log:        cfg.Log,
		UserBus:    userBus,
		HomeBus:    homeBus,
		AuthClient: cfg.AuthClient,
	})

	productapi.Routes(app, productapi.Config{
		Log:        cfg.Log,
		UserBus:    userBus,
		ProductBus: productBus,
		AuthClient: cfg.AuthClient,
	})

	vproductapi.Routes(app, vproductapi.Config{
		Log:         cfg.Log,
		UserBus:     userBus,
		VProductBus: vproductBus,
		AuthClient:  cfg.AuthClient,
	})
}
