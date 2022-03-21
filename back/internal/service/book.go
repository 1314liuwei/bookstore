package service

import (
	"back/internal/model"
	"back/internal/service/internal/dao"
	"back/internal/service/internal/do"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

type (
	sBook struct{}
)

var (
	insBook = sBook{}
)

func Book() *sBook {
	return &insBook
}

func (b sBook) Insert(ctx context.Context, in model.BookInsert) error {
	return dao.Books.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		_, err := dao.Users.Ctx(ctx).Data(do.Books{
			Name:        in.Name,
			Author:      in.Author,
			Ebook:       in.Ebook,
			Cover:       in.Cover,
			Price:       in.Price,
			Description: in.Description,
		}).Insert()
		if err != nil {
			return err
		}

		return nil
	})
}

func (b sBook) Query(ctx context.Context, in model.BookQueryInfo) (error, gdb.Result) {
	var query do.Books
	err := gconv.Scan(in, &query)
	if err != nil {
		return err, nil
	}

	all, err := dao.Books.Ctx(ctx).WhereOr(query).All()
	if err != nil {
		return err, nil
	}

	return nil, all
}
