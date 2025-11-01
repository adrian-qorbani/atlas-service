package checkapi

import (
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

func Routes(app *web.App) {

	app.HandleFuncNoMiddleware("GET /liveness", liveness)
	app.HandleFuncNoMiddleware("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testError)
	app.HandleFunc("GET /testpanic", testPanic)
}
