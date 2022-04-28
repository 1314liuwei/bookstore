// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-04-28 22:43:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrdersDao is the data access object for table orders.
type OrdersDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns OrdersColumns // columns contains all the column names of Table for convenient usage.
}

// OrdersColumns defines and stores column names for table orders.
type OrdersColumns struct {
	Id        string //
	Oid       string // 订单id号
	Status    string // 订单状态
	Ebook     string // 电子书地址
	CreatedAt string //
	UpdateAt  string //
	BookOrder string //
	UserOrder string //
}

//  ordersColumns holds the columns for table orders.
var ordersColumns = OrdersColumns{
	Id:        "id",
	Oid:       "oid",
	Status:    "status",
	Ebook:     "ebook",
	CreatedAt: "created_at",
	UpdateAt:  "update_at",
	BookOrder: "book_order",
	UserOrder: "user_order",
}

// NewOrdersDao creates and returns a new DAO object for table data access.
func NewOrdersDao() *OrdersDao {
	return &OrdersDao{
		group:   "default",
		table:   "orders",
		columns: ordersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrdersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrdersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrdersDao) Columns() OrdersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrdersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrdersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrdersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
