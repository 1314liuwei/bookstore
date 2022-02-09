package v1

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type BookQueryReq struct {
	g.Meta   `path:"/book/category" method:"POST" tags:"Book" summary:"Query books by category"`
	Category int64 `v:"required"`
}
type BookQueryRes struct {
	Books gdb.Result
}
