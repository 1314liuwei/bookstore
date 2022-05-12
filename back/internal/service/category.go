package service

import (
	"back/internal/consts"
	"back/internal/model/entity"
	"back/internal/service/internal/dao"
	"back/internal/service/internal/do"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
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

func (c sCategory) QueryAllCategories(ctx context.Context) ([]string, error) {
	all, err := dao.Categories.Ctx(ctx).Fields(dao.Categories.Columns().Name).All()
	if err != nil {
		return nil, err
	}

	var result []string
	result = append(result, consts.AllCagetory)
	for _, record := range all {
		result = append(result, gconv.String(record[dao.Categories.Columns().Name]))
	}

	return result, nil
}
