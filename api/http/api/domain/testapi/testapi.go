package testapi

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/adrian-qorbani/atlas-service/app/api/errs"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

type api struct{}

func newAPI() *api {
	return &api{}

}

func (api *api) testGood(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}

func (api *api) testError(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return errs.Newf(errs.FailedPrecondition, "this msg is trusted.")
	}
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}

func (api *api) testPanic(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		panic("PANICKING!!!!!")
	}
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}
