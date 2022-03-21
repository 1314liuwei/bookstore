// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-03-21 20:57:03
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GoadminUsers is the golang structure of table goadmin_users for DAO operations like Where/Data.
type GoadminUsers struct {
	g.Meta        `orm:"table:goadmin_users, do:true"`
	Id            interface{} //
	Username      interface{} //
	Password      interface{} //
	Name          interface{} //
	Avatar        interface{} //
	RememberToken interface{} //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
