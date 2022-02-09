package service

import (
	"back/internal/model"
	"back/internal/model/entity"
	"back/internal/service/internal/dao"
	"back/internal/service/internal/do"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
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

func (b sBook) QueryById(ctx context.Context, id int64) (error, *entity.Books) {
	var book *entity.Books

	err := dao.Books.Ctx(ctx).Where(do.Books{Id: id}).Scan(&book)
	if err != nil {
		return err, nil
	}

	if book == nil {
		return gerror.Newf(`The book with ID '%s' cannot be queried`, id), nil
	}

	return nil, book
}

func (b sBook) QueryByIds(ctx context.Context, ids []int64) (error, []*entity.Books) {
	var result []*entity.Books
	for _, id := range ids {
		err, e := b.QueryById(ctx, id)
		if err != nil {
			return err, result
		}

		result = append(result, e)
	}
	return nil, result
}
