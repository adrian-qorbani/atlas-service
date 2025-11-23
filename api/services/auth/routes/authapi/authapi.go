// Package authapi maintains the web based api for auth access.
package authapi

import (
	"context"
	"net/http"

	"github.com/adrian-qorbani/atlas-service/app/api/authclient"
	"github.com/adrian-qorbani/atlas-service/app/api/errs"
	"github.com/adrian-qorbani/atlas-service/app/api/middleware"
	"github.com/adrian-qorbani/atlas-service/business/api/auth"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

type api struct {
	auth *auth.Auth
}

func newAPI(ath *auth.Auth) *api {
	return &api{
		auth: ath,
	}
}

func (api *api) token(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	kid := web.Param(r, "kid")
	if kid == "" {
		return errs.Newf(errs.FailedPrecondition, "missing kid")
	}

	// The BearerBasic middleware function generates the claims.
	claims := middleware.GetClaims(ctx)

	tkn, err := api.auth.GenerateToken(kid, claims)
	if err != nil {
		return errs.New(errs.Internal, err)
	}

	token := struct {
		Token string `json:"token"`
	}{
		Token: tkn,
	}

	return web.Respond(ctx, w, token, http.StatusOK)
}

func (api *api) authenticate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	// The middleware is actually handling the authentication. So if the code
	// gets to this handler, authentication passed.

	userID, err := middleware.GetUserID(ctx)
	if err != nil {
		return errs.New(errs.Unauthenticated, err)
	}

	resp := authclient.AuthenticateResp{
		UserID: userID,
		Claims: middleware.GetClaims(ctx),
	}

	return web.Respond(ctx, w, resp, http.StatusOK)
}

func (api *api) authorize(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var auth authclient.Authorize
	if err := web.Decode(r, &auth); err != nil {
		return errs.New(errs.FailedPrecondition, err)
	}

	if err := api.auth.Authorize(ctx, auth.Claims, auth.UserID, auth.Rule); err != nil {
		return errs.Newf(errs.Unauthenticated, "authorize: you are not authorized for that action, claims[%v] rule[%v]: %s", auth.Claims.Roles, auth.Rule, err)
	}

	return web.Respond(ctx, w, nil, http.StatusNoContent)
}
