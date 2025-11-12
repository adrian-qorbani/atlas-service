package checkapi

import (
	"github.com/adrian-qorbani/atlas-service/api/services/api/middleware"
	"github.com/adrian-qorbani/atlas-service/business/api/auth"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

func Routes(app *web.App, a *auth.Auth) {

	authen := middleware.Authenticate(a)
	athAdminOnly := middleware.Authorize(a, auth.RuleAdminOnly)

	app.HandleFuncNoMiddleware("GET /liveness", liveness)
	app.HandleFuncNoMiddleware("GET /readiness", readiness)
	app.HandleFunc("GET /testerror", testError)
	app.HandleFunc("GET /testpanic", testPanic)
	app.HandleFunc("GET /testauth", liveness, authen, athAdminOnly)

}
