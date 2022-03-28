package v1

import "github.com/gogf/gf/v2/frame/g"

type CategoryQueryAllReq struct {
	g.Meta `path:"/category/all" method:"GET" tags:"Category" summary:"Query all categories"`
}
type CategoryQueryAllRes struct {
	Categories []string `json:"categories"`
}
