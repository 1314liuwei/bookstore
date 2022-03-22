package controller

import (
	v1 "back/api/v1"
	"back/internal/model"
	"back/internal/service"
	"context"
)

type cShoppingCart struct{}

var ShoppingCart = cShoppingCart{}

func (s cShoppingCart) Create(ctx context.Context, req *v1.ShoppingCartCreateReq) (res *v1.ShoppingCartCreateRes, err error) {
	id, err := service.ShoppingCarts().CreateShoppingCart(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.ShoppingCartCreateRes{
		SId: id,
	}
	return
}

func (s cShoppingCart) AddBooks(ctx context.Context, req *v1.ShoppingCartAddBooksReq) (res *v1.ShoppingCartAddBooksRes, err error) {
	err = service.ShoppingCarts().AddBook(ctx, model.ShoppingCartChangeBook{
		BookIds: req.BookIds,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (s cShoppingCart) RemoveBooks(ctx context.Context, req *v1.ShoppingCartRemoveBooksReq) (res *v1.ShoppingCartRemoveBooksRes, err error) {
	err = service.ShoppingCarts().RemoveBook(ctx, model.ShoppingCartChangeBook{
		BookIds: req.BookIds,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (s cShoppingCart) Empty(ctx context.Context, req *v1.ShoppingCartEmptyReq) (res *v1.ShoppingCartEmptyRes, err error) {
	err = service.ShoppingCarts().Empty(ctx, model.ShoppingCartChangeBook{})
	if err != nil {
		return nil, err
	}
	return
}

func (s cShoppingCart) Settle(ctx context.Context, req *v1.ShoppingCartEmptyReq) (res *v1.ShoppingCartEmptyRes, err error) {
	err = service.ShoppingCarts().Empty(ctx, model.ShoppingCartChangeBook{})
	if err != nil {
		return nil, err
	}
	return
}
