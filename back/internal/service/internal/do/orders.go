// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-04-28 22:43:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Orders is the golang structure of table orders for DAO operations like Where/Data.
type Orders struct {
	g.Meta    `orm:"table:orders, do:true"`
	Id        interface{} //
	Oid       interface{} // 订单id号
	Status    interface{} // 订单状态
	Ebook     interface{} // 电子书地址
	CreatedAt *gtime.Time //
	UpdateAt  *gtime.Time //
	BookOrder interface{} //
	UserOrder interface{} //
}
