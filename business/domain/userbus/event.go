package userbus

import (
	"encoding/json"
	"fmt"

	"github.com/adrian-qorbani/atlas-service/business/api/delegate"
	"github.com/google/uuid"
)

// DomainName represents the name of this domain.
const DomainName = "user"

// Set of delegate actions.
const (
	ActionDeleted = "deleted"
	ActionUpdated = "updated"
)

// ActionDeletedParms represents the parameters for the deleted action.
type ActionDeletedParms struct {
	UserID uuid.UUID
}

// ActionUpdatedParms represents the parameters for the updated action.
type ActionUpdatedParms struct {
	UserID uuid.UUID
	UpdateUser
}

// String returns a string representation of the action parameters.
func (act *ActionDeletedParms) String() string {
	return fmt.Sprintf("&EventParamsUpdated{UserID:%v}", act.UserID)
}

// Marshal returns the event parameters encoded as JSON.
func (act *ActionDeletedParms) Marshal() ([]byte, error) {
	return json.Marshal(act)
}

// ActionDeletedData constructs the data for the deleted action.
func ActionDeletedData(userID uuid.UUID) delegate.Data {
	params := ActionDeletedParms{
		UserID: userID,
	}

	rawParams, err := params.Marshal()
	if err != nil {
		panic(err)
	}

	return delegate.Data{
		Domain:    DomainName,
		Action:    ActionDeleted,
		RawParams: rawParams,
	}
}

// String returns a string representation of the action parameters.
func (au *ActionUpdatedParms) String() string {
	return fmt.Sprintf("&EventParamsUpdated{UserID:%v, Enabled:%v}", au.UserID, au.Enabled)
}

// Marshal returns the event parameters encoded as JSON.
func (au *ActionUpdatedParms) Marshal() ([]byte, error) {
	return json.Marshal(au)
}

// ActionUpdatedData constructs the data for the updated action.
func ActionUpdatedData(uu UpdateUser, userID uuid.UUID) delegate.Data {
	params := ActionUpdatedParms{
		UserID: userID,
		UpdateUser: UpdateUser{
			Enabled: uu.Enabled,
		},
	}

	rawParams, err := params.Marshal()
	if err != nil {
		panic(err)
	}

	return delegate.Data{
		Domain:    DomainName,
		Action:    ActionUpdated,
		RawParams: rawParams,
	}
}
