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
	"github.com/gogf/gf/v2/util/gconv"
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

func (o sOrder) Created(ctx context.Context, in model.Order) error {
	userId := Context().Get(ctx).User.Id
	if ok, err := o.IsUserExist(ctx, userId); !ok || err != nil {
		return gerror.New("User does not exist")
	}

	// 检查 book 是否有效
	for _, i := range in.Data {
		if ok, err := o.IsBookExist(ctx, int64(i)); !ok || err != nil {
			return gerror.New("Book does not exist")
		}
	}

	// oid 为当前时间戳
	var result []int64
	err := dao.Orders.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		temp := make(map[int]struct{})

		for _, id := range in.Data {
			if _, ok := temp[id]; ok {
				continue
			} else {
				temp[id] = struct{}{}
			}

			one, err := dao.Books.Ctx(ctx).Where(do.Books{Id: id}).One()
			if err != nil {
				return err
			}

			// 当前订单号
			oid := time.Now().UnixMilli()
			// 将订单号加入结果列表
			result = append(result, oid)
			_, err = dao.Orders.Ctx(ctx).Data(do.Orders{
				Oid:       oid,
				Status:    consts.OrderToBePaid,
				BookOrder: id,
				UserOrder: userId,
				Ebook:     one[dao.Books.Columns().Ebook],
			}).Insert()
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
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

func (o sOrder) QueryAll(ctx context.Context, in model.OrderQueryAll) ([]model.OrderQueryAllOutput, error) {
	all, err := dao.Orders.Ctx(ctx).Page(in.Page, consts.Limit).All()
	if err != nil {
		return nil, err
	}

	var data []model.OrderQueryAllOutput
	for _, record := range all {
		one, err := dao.Books.Ctx(ctx).Where(do.Books{Id: record[dao.Orders.Columns().BookOrder]}).One()
		if err != nil {
			return nil, err
		}

		temp := model.OrderQueryAllOutput{
			Price:     gconv.Float64(one[dao.Books.Columns().Price]),
			BookImage: gconv.String(one[dao.Books.Columns().Cover]),
			BookName:  gconv.String(one[dao.Books.Columns().Name]),
			BookID:    gconv.Int(one[dao.Books.Columns().Id]),

			CreatedAt: gconv.String(record[dao.Orders.Columns().CreatedAt]),
			Ebook:     gconv.String(record[dao.Orders.Columns().Ebook]),
			Id:        gconv.Int(record[dao.Orders.Columns().Id]),
			Oid:       gconv.Int64(record[dao.Orders.Columns().Oid]),
			Status:    gconv.String(record[dao.Orders.Columns().Status]),
			UpdateAt:  gconv.String(record[dao.Orders.Columns().UpdateAt]),
			UserOrder: gconv.Int(record[dao.Orders.Columns().UserOrder]),
		}
		data = append(data, temp)
	}
	return data, nil
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
