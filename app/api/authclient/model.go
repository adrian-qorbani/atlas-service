package authclient

import (
	"github.com/adrian-qorbani/atlas-service/business/api/auth"
	"github.com/google/uuid"
)

// Authorize defines the information required to perform an authorization.
type Authorize struct {
	UserID uuid.UUID
	Claims auth.Claims
	Rule   string
}

// AuthenticateResp defines the information that will be received on authenticate.
type AuthenticateResp struct {
	UserID uuid.UUID
	Claims auth.Claims
}
