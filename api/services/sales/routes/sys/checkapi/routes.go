package checkapi

import (
	"net/http"
)

func Routes(m *http.ServeMux) {

	m.HandleFunc("GET /liveness", liveness)
	m.HandleFunc("GET /liveness", readiness)
}
