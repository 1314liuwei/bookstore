package v1

import (
	"back/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type UserProfileReq struct {
	g.Meta `path:"/user/profile" method:"GET" tags:"User" summary:"Get the profile of current user"`
}
type UserProfileRes struct {
	*entity.Users
}

type UserSignUpReq struct {
	g.Meta    `path:"/user/sign_up" method:"POST" tags:"User" summary:"Sign up a new user account"`
	Username  string `v:"required"`
	Password  string `v:"required|password"`
	Password2 string `v:"required|password|same:Password"`
}
type UserSignUpRes struct{}

type UserSignInReq struct {
	g.Meta   `path:"/user/sign_in" method:"POST" tags:"User" summary:"Sign in with exists account"`
	Username string `v:"required"`
	Password string `v:"required|password"`
}
type UserSignInRes struct{}

type UserSignOutReq struct {
	g.Meta `path:"/user/sign_out" method:"POST" tags:"User" summary:"Sign out of account"`
}
type UserSignOutRes struct{}

type UserUpgradeTypeVIPReq struct {
	g.Meta `path:"/user/upgrade_vip" method:"POST" tags:"User" summary:"Upgrade user to VIP"`
}
type UserUpgradeTypeVIPRes struct{}
