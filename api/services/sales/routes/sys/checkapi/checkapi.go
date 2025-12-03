// package maintains the web based api for system access
package checkapi

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/adrian-qorbani/atlas-service/app/api/errs"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
	"github.com/go-json-experiment/json"
	"github.com/jmoiron/sqlx"
)

// temp
type status struct {
	Status string
}

type api struct {
	db *sqlx.DB
}

func newAPI(db *sqlx.DB) *api {
	return &api{
		db: db,
	}

}

func (s status) Encode() ([]byte, string, error) {
	data, err := json.Marshal(s)
	return data, "application/json", err
}

func (api *api) liveness(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	resp := status{Status: "OK"}
	return web.Respond(ctx, w, resp, http.StatusOK)
}

func (api *api) readiness(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	resp := status{Status: "OK"}
	return web.Respond(ctx, w, resp, http.StatusOK)
}

func (api *api) testError(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return errs.Newf(errs.FailedPrecondition, "this msg is trusted.")
	}
	resp := status{Status: "OK"}
	return web.Respond(ctx, w, resp, http.StatusOK)
}

func (api *api) testPanic(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		panic("PANICKING!!!!!")
	}
	resp := status{Status: "OK"}
	return web.Respond(ctx, w, resp, http.StatusOK)
}
