package middleware

import (
	"context"
	"errors"

	"github.com/adrian-qorbani/atlas-service/app/api/errs"
	"github.com/adrian-qorbani/atlas-service/business/api/auth"
)

// ErrInvalidID represents a condition where the id is not a uuid.
var ErrInvalidID = errors.New("ID is not in its proper form")

// Authorize validates authorization via the auth service.
func Authorize(ctx context.Context, auth *auth.Auth, rule string, handler Handler) error {
	userID, err := GetUserID(ctx)
	if err != nil {
		return errs.New(errs.Unauthenticated, err)
	}

	if err := auth.Authorize(ctx, GetClaims(ctx), userID, rule); err != nil {
		return errs.Newf(errs.Unauthenticated, "authorize")
	}

	return handler(ctx)
}
