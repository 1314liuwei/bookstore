// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT. Created at 2022-04-28 22:43:14
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure of table users for DAO operations like Where/Data.
type Users struct {
	g.Meta    `orm:"table:users, do:true"`
	Id        interface{} //
	Username  interface{} //
	Password  interface{} //
	Type      interface{} //
	CreatedAt *gtime.Time //
}
