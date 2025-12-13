package productapp

import (
	"github.com/adrian-qorbani/atlas-service/business/api/order"
	"github.com/adrian-qorbani/atlas-service/business/domain/productbus"
)

var defaultOrderBy = order.NewBy("product_id", order.ASC)

var orderByFields = map[string]string{
	"product_id": productbus.OrderByProductID,
	"name":       productbus.OrderByName,
	"cost":       productbus.OrderByCost,
	"quantity":   productbus.OrderByQuantity,
	"user_id":    productbus.OrderByUserID,
}
