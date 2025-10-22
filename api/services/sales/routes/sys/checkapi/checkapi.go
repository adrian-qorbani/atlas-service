// package maintains the web based api for system access
package checkapi

import (
	"encoding/json"
	"net/http"
)

func liveness(w http.ResponseWriter, req *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(status)
}

func readiness(w http.ResponseWriter, req *http.Request) {
	status := struct {
		Status string
	}{
		Status: "OK",
	}
	json.NewEncoder(w).Encode(status)
}
