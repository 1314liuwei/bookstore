// Code generated by entc, DO NOT EDIT.

package order

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeBook holds the string denoting the book edge name in mutations.
	EdgeBook = "book"
	// Table holds the table name of the order in the database.
	Table = "orders"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "orders"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_order"
	// BookTable is the table that holds the book relation/edge.
	BookTable = "orders"
	// BookInverseTable is the table name for the Book entity.
	// It exists in this package in order to avoid circular dependency with the "book" package.
	BookInverseTable = "books"
	// BookColumn is the table column denoting the book relation/edge.
	BookColumn = "book_order"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldAmount,
	FieldStatus,
	FieldCreatedAt,
	FieldUpdateAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "orders"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"book_order",
	"user_order",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdateAt holds the default value on creation for the "update_at" field.
	DefaultUpdateAt func() time.Time
	// UpdateDefaultUpdateAt holds the default value on update for the "update_at" field.
	UpdateDefaultUpdateAt func() time.Time
)

// Status defines the type for the "status" enum field.
type Status string

// Status values.
const (
	StatusCompleted    Status = "completed"
	StatusToBePaid     Status = "to_be_paid"
	StatusTransferring Status = "transferring"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusCompleted, StatusToBePaid, StatusTransferring:
		return nil
	default:
		return fmt.Errorf("order: invalid enum value for status field: %q", s)
	}
}
