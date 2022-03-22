package service

import (
	"back/internal/model"
	"back/internal/service/internal/dao"
	"back/internal/service/internal/do"
	"context"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type sShoppingCart struct{}

var insShoppingCart = sShoppingCart{}

func ShoppingCarts() *sShoppingCart {
	return &insShoppingCart
}

func (s sShoppingCart) CreateShoppingCart(ctx context.Context) (int64, error) {
	userId := Context().Get(ctx).User.Id
	if ok, err := s.IsUserExist(ctx, userId); !ok || err != nil {
		return 0, gerror.New("User does not exist")
	}

	oid := time.Now().UnixMilli()

	return oid, dao.Orders.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Orders.Ctx(ctx).Data(do.Orders{
			Oid:       oid,
			UserOrder: userId,
		}).Insert()
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *sShoppingCart) AddBook(ctx context.Context, in model.ShoppingCartChangeBook) error {
	userId := Context().Get(ctx).User.Id
	if ok, err := s.IsUserExist(ctx, userId); !ok || err != nil {
		return gerror.New("User does not exist")
	}

	return dao.ShoppingCarts.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		for _, id := range in.BookIds {
			insert, err := dao.ShoppingCarts.Ctx(ctx).Data(do.ShoppingCarts{
				BookShoppingCart: id,
				UserShoppingCart: userId,
			}).Insert()
			if err != nil {
				return err
			}

			affected, err := insert.RowsAffected()
			if err != nil || affected == 0 {
				return gerror.New("Insert failed")
			}
		}
		return nil
	})
}

func (s *sShoppingCart) RemoveBook(ctx context.Context, in model.ShoppingCartChangeBook) error {
	userId := Context().Get(ctx).User.Id
	if ok, err := s.IsUserExist(ctx, userId); !ok || err != nil {
		return gerror.New("User does not exist")
	}

	return dao.ShoppingCarts.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		for _, id := range in.BookIds {
			insert, err := dao.ShoppingCarts.Ctx(ctx).Data(do.ShoppingCarts{
				BookShoppingCart: id,
				UserShoppingCart: userId,
			}).Delete()
			if err != nil {
				return err
			}

			affected, err := insert.RowsAffected()
			if err != nil || affected == 0 {
				return gerror.New("Delete failed")
			}
		}
		return nil
	})
}

func (s *sShoppingCart) Empty(ctx context.Context, in model.ShoppingCartChangeBook) error {
	userId := Context().Get(ctx).User.Id
	if ok, err := s.IsUserExist(ctx, userId); !ok || err != nil {
		return gerror.New("User does not exist")
	}

	return dao.ShoppingCarts.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		insert, err := dao.ShoppingCarts.Ctx(ctx).Where(do.ShoppingCarts{
			UserShoppingCart: userId,
		}).Delete()
		if err != nil {
			return err
		}

		_, err = insert.RowsAffected()
		if err != nil {
			return gerror.New("Delete failed")
		}

		return nil
	})
}

func (s sShoppingCart) IsUserExist(ctx context.Context, id int64) (bool, error) {
	count, err := dao.Users.Ctx(ctx).Where(do.Users{
		Id: id,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 1, nil
}
