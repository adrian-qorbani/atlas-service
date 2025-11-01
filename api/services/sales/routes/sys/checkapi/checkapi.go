// package maintains the web based api for system access
package checkapi

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/adrian-qorbani/atlas-service/app/api/errs"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// temp
type status struct {
	Status string
}

func (s status) Encode() ([]byte, string, error) {
	data, err := json.Marshal(s)
	return data, "application/json", err
}

func liveness(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	resp := status{Status: "OK"}
	return web.Respond(ctx, w, resp, http.StatusOK)
}

func readiness(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	resp := status{Status: "OK"}
	return web.Respond(ctx, w, resp, http.StatusOK)
}

// temp test
func testError(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		return errs.Newf(errs.FailedPrecondition, "this msg is trusted.")
	}
	resp := status{Status: "OK"}
	return web.Respond(ctx, w, resp, http.StatusOK)
}

func testPanic(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	if n := rand.Intn(100); n%2 == 0 {
		panic("PANICKING!!!!!")
	}
	resp := status{Status: "OK"}
	return web.Respond(ctx, w, resp, http.StatusOK)
}
