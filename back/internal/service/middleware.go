package service

import (
	"back/internal/model"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sMiddleware struct{}
)

var (
	insMiddleware = sMiddleware{}
)

func Middleware() *sMiddleware {
	return &insMiddleware
}

func (m sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context().Init(r, customCtx)
	if user := Session().GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			Id:       user.Id,
			Username: user.Username,
			Type:     user.Type,
		}
	}

	r.Middleware.Next()
}

func (m sMiddleware) Auth(r *ghttp.Request) {
	// TODO: Auth
}

func (m sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
