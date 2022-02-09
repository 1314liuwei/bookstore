package service

import (
	"back/internal/consts"
	"back/internal/model"
	"back/internal/service/internal/dao"
	"back/internal/service/internal/do"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
)

type (
	sOrder struct{}
)

var (
	insOrder = sOrder{}
)

func Order() *sOrder {
	return &insOrder
}

func (o sOrder) Created(ctx context.Context, in model.OrderCreate) error {
	if in.Amount <= 0 {
		return gerror.New("Wrong amount")
	}

	userId := Context().Get(ctx).User.Id
	if ok, err := o.IsUserExist(ctx, userId); !ok || err != nil {
		return gerror.New("User does not exist")
	}

	if ok, err := o.IsBookExist(ctx, in.BookId); !ok || err != nil {
		return gerror.New("Book does not exist")
	}

	return dao.Orders.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Orders.Ctx(ctx).Data(do.Orders{
			Amount:    in.Amount,
			Status:    consts.OrderToBePaid,
			BookOrder: in.BookId,
			UserOrder: userId,
		}).Insert()
		if err != nil {
			return err
		}
		return nil
	})
}

func (o sOrder) IsUserExist(ctx context.Context, id int64) (bool, error) {
	count, err := dao.Users.Ctx(ctx).Where(do.Users{
		Id: id,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 1, nil
}

func (o sOrder) IsBookExist(ctx context.Context, id int64) (bool, error) {
	count, err := dao.Books.Ctx(ctx).Where(do.Users{
		Id: id,
	}).Count()
	if err != nil {
		return false, err
	}
	return count == 1, nil
}
