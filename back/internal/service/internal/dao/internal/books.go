// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-03-21 20:57:03
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BooksDao is the data access object for table books.
type BooksDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns BooksColumns // columns contains all the column names of Table for convenient usage.
}

// BooksColumns defines and stores column names for table books.
type BooksColumns struct {
	Id          string //
	Name        string //
	Author      string //
	Description string //
	Ebook       string //
	Cover       string //
	Price       string //
	CreatedAt   string //
	Category    string //
}

//  booksColumns holds the columns for table books.
var booksColumns = BooksColumns{
	Id:          "id",
	Name:        "name",
	Author:      "author",
	Description: "description",
	Ebook:       "ebook",
	Cover:       "cover",
	Price:       "price",
	CreatedAt:   "created_at",
	Category:    "category",
}

// NewBooksDao creates and returns a new DAO object for table data access.
func NewBooksDao() *BooksDao {
	return &BooksDao{
		group:   "default",
		table:   "books",
		columns: booksColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BooksDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BooksDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BooksDao) Columns() BooksColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BooksDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BooksDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BooksDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
