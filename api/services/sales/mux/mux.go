package mux

import (
	"os"

	"github.com/adrian-qorbani/atlas-service/api/services/api/middleware"
	"github.com/adrian-qorbani/atlas-service/api/services/sales/routes/sys/checkapi"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

func WebAPI(log *logger.Logger, shutdown chan os.Signal) *web.App {
	m := web.NewApp(shutdown, middleware.Logger(log), middleware.Errors(log), middleware.Panics())
	checkapi.Routes(m)

	return m
}
