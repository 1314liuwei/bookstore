package controller

import (
	v1 "back/api/v1"
	"back/internal/model"
	"back/internal/service"
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

type cEBook struct {
}

var EBook = cEBook{}

func (b cEBook) GetFile(ctx context.Context, req *v1.EBookGetFileReq) (res *v1.EBookGetFileRes, err error) {
	file, err := service.EBook().GetFile(ctx, model.EBookGetFile{File: req.File})
	if err != nil {
		return nil, err
	}
	ghttp.RequestFromCtx(ctx).Response.ServeFileDownload(file)
	return
}
