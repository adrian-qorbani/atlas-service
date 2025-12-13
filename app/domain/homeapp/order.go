package homeapp

import (
	"github.com/adrian-qorbani/atlas-service/business/api/order"
	"github.com/adrian-qorbani/atlas-service/business/domain/homebus"
)

var defaultOrderBy = order.NewBy("home_id", order.ASC)

var orderByFields = map[string]string{
	"home_id": homebus.OrderByID,
	"type":    homebus.OrderByType,
	"user_id": homebus.OrderByUserID,
}
