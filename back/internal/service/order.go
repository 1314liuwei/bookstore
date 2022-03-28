package service

import (
	"back/internal/consts"
	"back/internal/model"
	"back/internal/service/internal/dao"
	"back/internal/service/internal/do"
	"context"
	"time"

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

func (o sOrder) Created(ctx context.Context, in model.Order) ([]int64, error) {
	userId := Context().Get(ctx).User.Id
	if ok, err := o.IsUserExist(ctx, userId); !ok || err != nil {
		return nil, gerror.New("User does not exist")
	}

	// 检查 book 是否有效
	for _, i := range in.Data {
		if ok, err := o.IsBookExist(ctx, int64(i.BookId)); !ok || err != nil {
			return nil, gerror.New("Book does not exist")
		}
	}

	// oid 为当前时间戳
	var result []int64
	err := dao.Orders.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		temp := make(map[int]struct{})

		for _, data := range in.Data {
			if _, ok := temp[data.BookId]; ok {
				continue
			} else {
				temp[data.BookId] = struct{}{}
			}

			// 当前订单号
			oid := time.Now().UnixMilli()
			// 将订单号加入结果列表
			result = append(result, oid)
			_, err := dao.Orders.Ctx(ctx).Data(do.Orders{
				Oid:       oid,
				Status:    consts.OrderToBePaid,
				BookOrder: data.BookId,
				UserOrder: userId,
			}).Insert()
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (o sOrder) Remove(ctx context.Context, in model.OrderRemove) error {
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

func (o sOrder) UpdateStatus(ctx context.Context, in model.OrderUpdate) error {
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

func (o sOrder) Query(ctx context.Context, in model.OrderQuery) (gdb.Result, error) {
	all, err := dao.Orders.Ctx(ctx).Where(do.Orders{
		Oid: in.OId,
	}).All()
	if err != nil {
		return nil, err
	}

	return all, nil
}

func (o sOrder) QueryAll(ctx context.Context, in model.OrderQueryAll) (gdb.Result, error) {
	all, err := dao.Orders.Ctx(ctx).Page(in.Page, consts.Limit).All()
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
