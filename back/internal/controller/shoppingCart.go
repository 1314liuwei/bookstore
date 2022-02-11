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
		ShoppingCartId: req.SId,
		BookIds:        req.BookIds,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (s cShoppingCart) RemoveBooks(ctx context.Context, req *v1.ShoppingCartRemoveBooksReq) (res *v1.ShoppingCartRemoveBooksRes, err error) {
	err = service.ShoppingCarts().RemoveBook(ctx, model.ShoppingCartChangeBook{
		ShoppingCartId: req.SId,
		BookIds:        req.BookIds,
	})
	if err != nil {
		return nil, err
	}
	return
}

func (s cShoppingCart) Remove(ctx context.Context, req *v1.ShoppingCartRemoveReq) (res *v1.ShoppingCartRemoveRes, err error) {
	err = service.ShoppingCarts().RemoveShoppingCart(ctx, model.ShoppingCartChangeBook{
		ShoppingCartId: req.SId,
	})
	if err != nil {
		return nil, err
	}
	return
}
