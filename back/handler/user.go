package handler

import (
	"back/model"
	"back/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Login(ctx *gin.Context) {
	var err error

	validate := validator.New()

	// 验证用户名
	username := ctx.PostForm("username")
	err = validate.Var(username, "email, required")
	if err != nil {
		return
	}

	passwd := ctx.PostForm("passwd")
	err = validate.Var(passwd, "alphanum")
	if err != nil {
		return
	}

	user, err := model.QueryUser(ctx, repository.Client, username, passwd)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"code": 200,
		"data": user,
	})
}

func SignUp(ctx *gin.Context) {
	var err error

	validate := validator.New()

	// 验证用户名
	username := ctx.PostForm("username")
	err = validate.Var(username, "email, required")
	if err != nil {
		return
	}

	passwd := ctx.PostForm("passwd")
	err = validate.Var(passwd, "alphanum")
	if err != nil {
		return
	}

	user, err := model.CreateUser(ctx, username, passwd, "normal")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "用户创建失败",
			"code": 503,
			"data": nil,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "用户创建成功",
			"code": 200,
			"data": user,
		})
	}
}
