// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"back/internal/service/internal/dao/internal"
)

// goadminOperationLogDao is the data access object for table goadmin_operation_log.
// You can define custom methods on it to extend its functionality as you wish.
type goadminOperationLogDao struct {
	*internal.GoadminOperationLogDao
}

var (
	// GoadminOperationLog is globally public accessible object for table goadmin_operation_log operations.
	GoadminOperationLog = goadminOperationLogDao{
		internal.NewGoadminOperationLogDao(),
	}
)

// Fill with you ideas below.
