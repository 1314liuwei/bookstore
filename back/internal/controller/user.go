package controller

import (
	"back/api/v1"
	"back/internal/model"
	"back/internal/service"
	"context"
)

type cUser struct{}

var User = cUser{}

func (u cUser) SignUp(ctx context.Context, req *v1.UserSignUpReq) (res *v1.UserSignUpRes, err error) {
	err = service.User().SignUp(ctx, model.UserInput{
		Username: req.Username,
		Password: req.Password,
	})
	return
}

func (u cUser) SignIn(ctx context.Context, req *v1.UserSignInReq) (res *v1.UserSignInRes, err error) {
	token, err := service.User().SignIn(ctx, model.UserInput{
		Username: req.Username,
		Password: req.Password,
	})

	res = &v1.UserSignInRes{
		Token: token,
	}
	return
}

func (u cUser) SignOut(ctx context.Context, req *v1.UserSignOutReq) (res *v1.UserSignOutRes, err error) {
	err = service.User().SignOut(ctx)
	return
}

func (u cUser) Profile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	res = &v1.UserProfileRes{
		Users: service.User().GetProfile(ctx),
	}
	return
}

func (u cUser) UpgradeVIP(ctx context.Context, req *v1.UserUpgradeTypeVIPReq) (res *v1.UserUpgradeTypeVIPRes, err error) {
	err = service.User().UpgradeVIP(ctx)
	return
}
