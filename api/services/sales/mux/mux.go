package mux

import (
	"net/http"

	"github.com/adrian-qorbani/atlas-service/api/services/sales/routes/sys/checkapi"
)

func WebAPI() *http.ServeMux {
	m := http.NewServeMux()

	checkapi.Routes(m)

	return m
}
