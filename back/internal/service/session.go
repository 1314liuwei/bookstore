package service

import (
	"back/internal/consts"
	"back/internal/model/entity"
	"context"
)

type (
	sSession struct{}
)

var (
	insSession = sSession{}
)

func Session() *sSession {
	return &insSession
}

func (s sSession) SetUser(ctx context.Context, user *entity.Users) error {
	return Context().Get(ctx).Session.Set(consts.UserSessionKey, user)
}

func (s sSession) GetUser(ctx context.Context) *entity.Users {
	customCtx := Context().Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.MustGet(consts.UserSessionKey); !v.IsNil() {
			var user *entity.Users
			err := v.Struct(&user)
			if err != nil {
				return nil
			}
			return user
		}
	}
	return nil
}

func (s sSession) RemoveUser(ctx context.Context) error {
	customCtx := Context().Get(ctx)
	if customCtx != nil {
		return customCtx.Session.Remove(consts.UserSessionKey)
	}
	return nil
}
