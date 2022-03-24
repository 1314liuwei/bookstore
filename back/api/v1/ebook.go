package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type EBookGetFileReq struct {
	g.Meta `path:"/ebook/:file" method:"GET" tags:"EBook" summary:"Get an ebook file"`
	File   string `path:"file" v:"required"`
}
type EBookGetFileRes struct {
}
