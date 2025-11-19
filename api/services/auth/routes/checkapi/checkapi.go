// package maintains the web based api for system access
package checkapi

import (
	"context"
	"net/http"

	"github.com/adrian-qorbani/atlas-service/foundation/web"
	"github.com/go-json-experiment/json"
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
