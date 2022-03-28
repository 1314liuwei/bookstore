package controller

import (
	v1 "back/api/v1"
	"back/internal/model"
	"back/internal/service"
	"context"
)

type cBook struct{}

var Book = cBook{}

func (b cBook) QueryBooksByCategory(ctx context.Context, req *v1.BookQueryByCategoryReq) (res *v1.BookQueryByCategoryRes, err error) {
	err, g := service.Book().Query(ctx, model.BookQueryInfo{
		Page:     req.Page,
		Category: req.Category,
	})
	if err != nil {
		return nil, err
	}

	res = &v1.BookQueryByCategoryRes{
		Books: g,
	}
	return
}

func (b cBook) QueryBooksByID(ctx context.Context, req *v1.BookQueryByIDReq) (res *v1.BookQueryByIDRes, err error) {
	err, g := service.Book().Query(ctx, model.BookQueryInfo{
		ID: req.ID,
	})
	if err != nil {
		return nil, err
	}

	res = &v1.BookQueryByIDRes{
		Books: g,
	}
	return
}

func (b cBook) QueryBooksRandom20(ctx context.Context, req *v1.BookQueryRandom20Req) (res *v1.BookQueryByCategoryRes, err error) {
	err, g := service.Book().QueryRandom20(ctx, model.BookQueryInfo{})
	if err != nil {
		return nil, err
	}

	res = &v1.BookQueryByCategoryRes{
		Books: g,
	}
	return
}
