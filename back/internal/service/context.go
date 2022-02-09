package service

import (
	"back/internal/consts"
	"back/internal/model"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sContext struct{}
)

var (
	insContext = sContext{}
)

func Context() *sContext {
	return &insContext
}

func (c sContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

func (c sContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.ContextKey)
	if value == nil {
		return nil
	}

	if m, ok := value.(*model.Context); ok {
		return m
	}
	return nil
}

func (c sContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	c.Get(ctx).User = ctxUser
}
