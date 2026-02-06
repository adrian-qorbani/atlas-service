// Package tranapi maintains the web based api for tran access.
package tranapi

import (
	"context"
	"net/http"

	"github.com/adrian-qorbani/atlas-service/app/api/errs"
	tranapp "github.com/adrian-qorbani/atlas-service/app/domain/transapp"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

type api struct {
	tranApp *tranapp.App
}

func newAPI(tranApp *tranapp.App) *api {
	return &api{
		tranApp: tranApp,
	}
}

func (api *api) create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app tranapp.NewTran
	if err := web.Decode(r, &app); err != nil {
		return errs.New(errs.FailedPrecondition, err)
	}

	prd, err := api.tranApp.Create(ctx, app)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, prd, http.StatusCreated)
}
