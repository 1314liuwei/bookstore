// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-03-21 20:57:03
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GoadminUserPermissionsDao is the data access object for table goadmin_user_permissions.
type GoadminUserPermissionsDao struct {
	table   string                        // table is the underlying table name of the DAO.
	group   string                        // group is the database configuration group name of current DAO.
	columns GoadminUserPermissionsColumns // columns contains all the column names of Table for convenient usage.
}

// GoadminUserPermissionsColumns defines and stores column names for table goadmin_user_permissions.
type GoadminUserPermissionsColumns struct {
	UserId       string //
	PermissionId string //
	CreatedAt    string //
	UpdatedAt    string //
}

//  goadminUserPermissionsColumns holds the columns for table goadmin_user_permissions.
var goadminUserPermissionsColumns = GoadminUserPermissionsColumns{
	UserId:       "user_id",
	PermissionId: "permission_id",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewGoadminUserPermissionsDao creates and returns a new DAO object for table data access.
func NewGoadminUserPermissionsDao() *GoadminUserPermissionsDao {
	return &GoadminUserPermissionsDao{
		group:   "default",
		table:   "goadmin_user_permissions",
		columns: goadminUserPermissionsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GoadminUserPermissionsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GoadminUserPermissionsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GoadminUserPermissionsDao) Columns() GoadminUserPermissionsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *GoadminUserPermissionsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GoadminUserPermissionsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GoadminUserPermissionsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
