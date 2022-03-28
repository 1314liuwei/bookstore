package controller

import (
	v1 "back/api/v1"
	"back/internal/service"
	"context"
)

type cCategory struct {
}

var Category = cCategory{}

func (c cCategory) QueryAllCategories(ctx context.Context, req *v1.CategoryQueryAllReq) (res *v1.CategoryQueryAllRes, err error) {
	categories, err := service.Category().QueryAllCategories(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.CategoryQueryAllRes{Categories: categories}
	return
}
