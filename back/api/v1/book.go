package v1

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type BookQueryByCategoryReq struct {
	g.Meta   `path:"/book/category" method:"POST" tags:"Book" summary:"Query books by category"`
	Category string `default:"全部"`
	Page     int    `default:"1"`
}
type BookQueryByCategoryRes struct {
	Books gdb.Result
}

type BookQueryByIDReq struct {
	g.Meta `path:"/book/id" method:"POST" tags:"Book" summary:"Query books by category"`
	ID     int `v:"required"`
}
type BookQueryByIDRes struct {
	Books gdb.Result
}

type BookQueryRandom20Req struct {
	g.Meta `path:"/book/20" method:"GET" tags:"Book" summary:"Query books by category"`
}
type BookQueryRandom20Res struct {
	Books gdb.Result
}
