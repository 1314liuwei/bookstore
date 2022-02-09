package controller

import (
	v1 "back/api/v1"
	"back/internal/consts"
	"back/internal/model"
	"back/internal/service"
	"context"
)

type cOrder struct{}

var Order = cOrder{}

func (o cOrder) Create(ctx context.Context, req *v1.OrderCreateReq) (res *v1.OrderCreateRes, err error) {
	err = service.Order().Created(ctx, model.Order{
		Amount: req.Amount,
		BookId: req.BookId,
	})
	return
}

func (o cOrder) Remove(ctx context.Context, req *v1.OrderRemoveReq) (res *v1.OrderRemoveRes, err error) {
	err = service.Order().Remove(ctx, model.Order{Id: req.Id})
	return
}

func (o cOrder) UpdateStatusCompleted(ctx context.Context, req *v1.OrderUpdateStatusCompletedReq) (res *v1.OrderUpdateStatusCompletedRes, err error) {
	err = service.Order().UpdateStatus(ctx, model.Order{Id: req.Id, Status: consts.OrderCompleted})
	return
}
