package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type EBookGetFileReq struct {
	g.Meta `path:"/ebook" method:"GET" tags:"EBook" summary:"Get an ebook file"`
	File   string `p:"file" v:"required"`
}
type EBookGetFileRes struct {
}
