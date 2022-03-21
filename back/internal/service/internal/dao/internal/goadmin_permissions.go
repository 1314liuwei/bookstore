// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-03-21 20:57:03
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoadminPermissionsDao is the data access object for table goadmin_permissions.
type GoadminPermissionsDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns GoadminPermissionsColumns // columns contains all the column names of Table for convenient usage.
}

// GoadminPermissionsColumns defines and stores column names for table goadmin_permissions.
type GoadminPermissionsColumns struct {
	Id         string //
	Name       string //
	Slug       string //
	HttpMethod string //
	HttpPath   string //
	CreatedAt  string //
	UpdatedAt  string //
}

//  goadminPermissionsColumns holds the columns for table goadmin_permissions.
var goadminPermissionsColumns = GoadminPermissionsColumns{
	Id:         "id",
	Name:       "name",
	Slug:       "slug",
	HttpMethod: "http_method",
	HttpPath:   "http_path",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewGoadminPermissionsDao creates and returns a new DAO object for table data access.
func NewGoadminPermissionsDao() *GoadminPermissionsDao {
	return &GoadminPermissionsDao{
		group:   "default",
		table:   "goadmin_permissions",
		columns: goadminPermissionsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoadminPermissionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoadminPermissionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoadminPermissionsDao) Columns() GoadminPermissionsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoadminPermissionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoadminPermissionsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoadminPermissionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
