package service

import (
	"back/internal/consts"
	"back/internal/model"
	"back/internal/service/internal/dao"
	"back/internal/service/internal/do"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"time"
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

func (o sOrder) Created(ctx context.Context, in model.Order) (int64, error) {
	if in.Amount <= 0 {
		return 0, gerror.New("Wrong amount")
	}

	userId := Context().Get(ctx).User.Id
	if ok, err := o.IsUserExist(ctx, userId); !ok || err != nil {
		return 0, gerror.New("User does not exist")
	}

	for _, i := range in.BookIds {
		if ok, err := o.IsBookExist(ctx, i); !ok || err != nil {
			return 0, gerror.New("Book does not exist")
		}
	}

	oid := time.Now().UnixMilli()

	return oid, dao.Orders.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		for _, id := range in.BookIds {
			_, err := dao.Orders.Ctx(ctx).Data(do.Orders{
				Oid:       oid,
				Amount:    in.Amount,
				Status:    consts.OrderToBePaid,
				BookOrder: id,
				UserOrder: userId,
			}).Insert()
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (o sOrder) Remove(ctx context.Context, in model.Order) error {
	result, err := dao.Orders.Ctx(ctx).Where(do.Orders{
		Oid: in.OId,
	}).Delete()
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return gerror.Newf(`Empty order with id: '%d' failed.`, in.OId)
	}
	return err
}

func (o sOrder) UpdateStatus(ctx context.Context, in model.Order) error {
	result, err := dao.Orders.Ctx(ctx).Where(do.Orders{
		Oid: in.OId,
	}).Update(do.Orders{
		Status: in.Status,
	})
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil || affected == 0 {
		return gerror.Newf(`Update order with id: '%d' failed.`, in.OId)
	}
	return err
}

func (o sOrder) Query(ctx context.Context, in model.Order) (gdb.Result, error) {
	all, err := dao.Orders.Ctx(ctx).Where(do.Orders{
		Oid: in.OId,
	}).All()
	if err != nil {
		return nil, err
	}

	return all, nil
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
