package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/adrian-qorbani/atlas-service/app/api/auth"
	"github.com/adrian-qorbani/atlas-service/app/api/authclient"
	"github.com/adrian-qorbani/atlas-service/app/api/errs"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// Authenticate validates authentication via the auth service.
func Authenticate(ctx context.Context, log *logger.Logger, client *authclient.Client, authorization string, handler Handler) error {
	resp, err := client.Authenticate(ctx, authorization)
	if err != nil {
		return errs.New(errs.Unauthenticated, err)
	}

	ctx = setUserID(ctx, resp.UserID)
	ctx = setClaims(ctx, resp.Claims)

	return handler(ctx)
}

// Bearer processes JWT authentication logic.
func Bearer(ctx context.Context, ath *auth.Auth, authorization string, handler Handler) error {
	claims, err := ath.Authenticate(ctx, authorization)
	if err != nil {
		return errs.New(errs.Unauthenticated, err)
	}

	if claims.Subject == "" {
		return errs.Newf(errs.Unauthenticated, "authorize: you are not authorized for that action, no claims")
	}

	subjectID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return errs.New(errs.Unauthenticated, fmt.Errorf("parsing subject: %w", err))
	}

	ctx = setUserID(ctx, subjectID)
	ctx = setClaims(ctx, claims)

	return handler(ctx)
}

// Basic processes basic authentication logic.
func Basic(ctx context.Context, handler Handler) error {
	claims := auth.Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   "3152420d-c05e-419e-a42a-7e5eb1151cac", //-to-do
			Issuer:    "service project",
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(8760 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		},
		Roles: []string{"ADMIN"},
	}

	subjectID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return errs.Newf(errs.Unauthenticated, "parsing subject: %s", err)
	}

	ctx = setUserID(ctx, subjectID)
	ctx = setClaims(ctx, claims)

	return handler(ctx)
}

// func parseBasicAuth(auth string) (string, string, bool) {
// 	parts := strings.Split(auth, " ")
// 	if len(parts) != 2 || parts[0] != "Basic" {
// 		return "", "", false
// 	}

// 	c, err := base64.StdEncoding.DecodeString(parts[1])
// 	if err != nil {
// 		return "", "", false
// 	}

// 	username, password, ok := strings.Cut(string(c), ":")
// 	if !ok {
// 		return "", "", false
// 	}

// 	return username, password, true
// }

// func processJWT(ctx context.Context, auth *auth.Auth, token string) (context.Context, error) {
// 	claims, err := auth.Authenticate(ctx, token)
// 	if err != nil {
// 		return ctx, errs.New(errs.Unauthenticated, err)
// 	}

// 	if claims.Subject == "" {
// 		return ctx, errs.Newf(errs.Unauthenticated, "no claims: youre not authorized for this action.")
// 	}

// 	subjectID, err := uuid.Parse(claims.Subject)
// 	if err != nil {
// 		return ctx, errs.New(errs.Unauthenticated, fmt.Errorf("parsing subject: %w", err))
// 	}

// 	ctx = setUserID(ctx, subjectID)
// 	ctx = setClaims(ctx, claims)

// 	return ctx, nil
// }
