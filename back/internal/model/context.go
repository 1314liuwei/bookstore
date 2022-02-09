package model

import "github.com/gogf/gf/v2/net/ghttp"

type Context struct {
	Session *ghttp.Session
	User    *ContextUser
}

type ContextUser struct {
	Id       int64
	Username string
	Type     string
}
