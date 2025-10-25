package mux

import (
	"os"

	"github.com/adrian-qorbani/atlas-service/api/services/sales/routes/sys/checkapi"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

func WebAPI(shutdown chan os.Signal) *web.App {
	m := web.NewApp(shutdown)
	checkapi.Routes(m)

	return m
}
