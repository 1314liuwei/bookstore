package service

import (
	"back/internal/consts"
	"back/internal/model"
	"back/internal/model/entity"
	"back/internal/service/internal/dao"
	"back/internal/service/internal/do"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	sUser struct{}
)

var (
	insUser = sUser{}
)

func User() *sUser {
	return &insUser
}

// SignUp User
func (u *sUser) SignUp(ctx context.Context, in model.UserInput) (err error) {
	// 用户名和密码不能为空
	if in.Username == "" || in.Password == "" {
		return gerror.New(`Username and Password cannot be empty`)
	}

	// 判断用户名是否重复
	available, err := u.IsUsernameAvailable(ctx, in.Username)
	if err != nil {
		return err
	}

	if !available {
		return gerror.Newf(`Username '%s' is not available`, in.Username)
	}

	// 返回结果
	return dao.Users.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Users.Ctx(ctx).Data(do.Users{
			Username: in.Username,
			Password: in.Password,
			Type:     consts.UserTypeNormal,
		}).Insert()
		return err
	})
}

func (u sUser) SignIn(ctx context.Context, in model.UserInput) (err error) {
	var user *entity.Users
	err = dao.Users.Ctx(ctx).Where(do.Users{
		Username: in.Username,
		Password: in.Password,
	}).Scan(&user)
	if err != nil {
		return err
	}

	if user == nil {
		return gerror.New(`Password not correct`)
	}

	if err = Session().SetUser(ctx, user); err != nil {
		return err
	}

	Context().SetUser(ctx, &model.ContextUser{
		Id:       user.Id,
		Username: user.Username,
		Type:     user.Type,
	})
	return
}

func (u sUser) SignOut(ctx context.Context) error {
	return Session().RemoveUser(ctx)
}

func (u sUser) GetProfile(ctx context.Context) *entity.Users {
	return Session().GetUser(ctx)
}

func (u sUser) UpgradeVIP(ctx context.Context) (err error) {
	var user *entity.Users
	user = Session().GetUser(ctx)
	if user == nil {
		return gerror.New(`Session error`)
	}

	return dao.Users.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Users.Ctx(ctx).Where(do.Users{
			Id: user.Id,
		}).Update(do.Users{
			Id:   user.Id,
			Type: consts.UserTypeVIP,
		})
		return err
	})
}

func (u sUser) IsUsernameAvailable(ctx context.Context, username string) (bool, error) {
	count, err := dao.Users.Ctx(ctx).Where(do.Users{
		Username: username,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 0, nil
}
