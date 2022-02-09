package v1

import (
	"back/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type OrderCreateReq struct {
	g.Meta `path:"/order/create" method:"POST" tags:"Order" summary:"Create an order"`
	Amount float64 `v:"required"`
	BookId int64   `v:"required"`
}
type OrderCreateRes struct {
	*entity.Orders
}
