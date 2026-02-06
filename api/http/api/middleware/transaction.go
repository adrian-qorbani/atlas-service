package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adrian-qorbani/atlas-service/app/api/middleware"
	"github.com/adrian-qorbani/atlas-service/business/api/transaction"
	"github.com/adrian-qorbani/atlas-service/foundation/logger"
	"github.com/adrian-qorbani/atlas-service/foundation/web"
)

// BeginCommitRollback executes the transaction middleware functionality.
func BeginCommitRollback(log *logger.Logger, bgn transaction.Beginner) web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			hdl := func(ctx context.Context) error {
				if err := handler(ctx, w, r); err != nil {
					return fmt.Errorf("EXECUTE TRANSACTION: %w", err)
				}
				return nil
			}

			return middleware.BeginCommitRollback(ctx, log, bgn, hdl)
		}

		return h
	}

	return m
}
