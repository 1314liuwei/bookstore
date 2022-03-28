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
	oid, err := service.Order().Created(ctx, model.Order{
		Data: req.Data,
	})
	res = &v1.OrderCreateRes{OIds: oid}
	return
}

func (o cOrder) Remove(ctx context.Context, req *v1.OrderRemoveReq) (res *v1.OrderRemoveRes, err error) {
	err = service.Order().Remove(ctx, model.OrderRemove{OId: req.Id})
	return
}

func (o cOrder) UpdateStatusCompleted(ctx context.Context, req *v1.OrderUpdateStatusCompletedReq) (res *v1.OrderUpdateStatusCompletedRes, err error) {
	err = service.Order().UpdateStatus(ctx, model.OrderUpdate{OId: req.Id, Status: consts.OrderCompleted})
	return
}

func (o cOrder) QueryAll(ctx context.Context, req *v1.OrderQueryAllReq) (res *v1.OrderQueryAllRes, err error) {
	all, err := service.Order().QueryAll(ctx, model.OrderQueryAll{Page: req.Page})
	if err != nil {
		return nil, err
	}
	res = &v1.OrderQueryAllRes{Data: all}
	return
}
