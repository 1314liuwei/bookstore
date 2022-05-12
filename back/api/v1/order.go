package v1

import (
	"back/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type OrderCreateReq struct {
	g.Meta `path:"/order/create" method:"POST" tags:"Order" summary:"Create an order"`
	Data   []int `v:"required"`
}
type OrderCreateRes struct {
	OIds []int64 `v:"required" json:"order_id"`
}

type OrderRemoveReq struct {
	g.Meta `path:"/order/remove" method:"POST" tags:"Order" summary:"Empty an order"`
	Id     int64 `v:"required"`
}
type OrderRemoveRes struct{}

type OrderUpdateStatusCompletedReq struct {
	g.Meta `path:"/order/completed" method:"POST" tags:"Order" summary:"Change order status to 'complete'"`
	Id     int64 `v:"required"`
}
type OrderUpdateStatusCompletedRes struct{}

type OrderQueryAllReq struct {
	g.Meta `path:"/order/all" method:"GET" tags:"Order" summary:"Query all orders"`
	Page   int `default:"1"`
}
type OrderQueryAllRes struct {
	Data []model.OrderQueryAllOutput
}
