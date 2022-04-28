// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-04-28 22:43:14
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShoppingCartsDao is the data access object for table shopping_carts.
type ShoppingCartsDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns ShoppingCartsColumns // columns contains all the column names of Table for convenient usage.
}

// ShoppingCartsColumns defines and stores column names for table shopping_carts.
type ShoppingCartsColumns struct {
	Id               string //
	Sid              string // 购物车id号
	CreatedAt        string //
	UpdateAt         string //
	BookShoppingCart string //
	UserShoppingCart string //
}

//  shoppingCartsColumns holds the columns for table shopping_carts.
var shoppingCartsColumns = ShoppingCartsColumns{
	Id:               "id",
	Sid:              "sid",
	CreatedAt:        "created_at",
	UpdateAt:         "update_at",
	BookShoppingCart: "book_shopping_cart",
	UserShoppingCart: "user_shopping_cart",
}

// NewShoppingCartsDao creates and returns a new DAO object for table data access.
func NewShoppingCartsDao() *ShoppingCartsDao {
	return &ShoppingCartsDao{
		group:   "default",
		table:   "shopping_carts",
		columns: shoppingCartsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ShoppingCartsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ShoppingCartsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ShoppingCartsDao) Columns() ShoppingCartsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ShoppingCartsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ShoppingCartsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ShoppingCartsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
