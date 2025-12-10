package userapp

import (
	"github.com/adrian-qorbani/atlas-service/business/api/order"
	"github.com/adrian-qorbani/atlas-service/business/domain/userbus"
)

var defaultOrderBy = order.NewBy("user_id", order.ASC)

var orderByFields = map[string]string{
	"user_id": userbus.OrderByID,
	"name":    userbus.OrderByName,
	"email":   userbus.OrderByEmail,
	"roles":   userbus.OrderByRoles,
	"enabled": userbus.OrderByEnabled,
}
