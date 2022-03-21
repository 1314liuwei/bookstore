// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-03-21 20:57:03
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GoadminSite is the golang structure of table goadmin_site for DAO operations like Where/Data.
type GoadminSite struct {
	g.Meta      `orm:"table:goadmin_site, do:true"`
	Id          interface{} //
	Key         interface{} //
	Value       interface{} //
	Description interface{} //
	State       interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
