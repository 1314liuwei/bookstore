package controller

import (
	v1 "back/api/v1"
	"back/internal/model"
	"back/internal/service"
	"context"
)

type cOrder struct{}

var Order = cOrder{}

func (o cOrder) Create(ctx context.Context, req *v1.OrderCreateReq) (res *v1.OrderCreateRes, err error) {
	err = service.Order().Created(ctx, model.OrderCreate{
		Amount: req.Amount,
		BookId: req.BookId,
	})
	return
}
