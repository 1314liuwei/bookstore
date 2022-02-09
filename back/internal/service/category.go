package service

import (
	"back/internal/model/entity"
	"back/internal/service/internal/dao"
	"back/internal/service/internal/do"
	"context"
)

type sCategory struct{}

var insCategory = sCategory{}

func Category() *sCategory {
	return &insCategory
}

func (c sCategory) IsCategoryExistByName(ctx context.Context, name string) (error, *entity.Categories) {
	var category *entity.Categories

	err := dao.Categories.Ctx(ctx).Where(do.Categories{
		Name: name,
	}).Scan(&category)

	return err, category
}

func (c sCategory) IsCategoryExistById(ctx context.Context, id int64) (error, *entity.Categories) {
	var category *entity.Categories

	err := dao.Categories.Ctx(ctx).Where(do.Categories{
		Id: id,
	}).Scan(&category)

	return err, category
}
