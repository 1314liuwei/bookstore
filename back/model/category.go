package model

import (
	"back/repository"
	"back/repository/ent/category"
	"context"
)

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateCategory(ctx context.Context, name string) (Category, error) {
	client := repository.GetDBClient()

	c, err := client.Category.Create().SetName(name).Save(ctx)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: c.ID, Name: c.Name}, nil
}

func QueryCategory(ctx context.Context, ID int) (Category, error) {
	client := repository.GetDBClient()
	c, err := client.Category.Query().Where(category.ID(ID)).Only(ctx)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: c.ID, Name: c.Name}, nil
}

func UpdateCategory(ctx context.Context, ID int, name string) (Category, error) {
	client := repository.GetDBClient()

	c, err := client.Category.UpdateOneID(ID).SetName(name).Save(ctx)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: c.ID, Name: c.Name}, nil
}

func DeleteCategory(ctx context.Context, ID int) error {
	client := repository.GetDBClient()
	_, err := client.Category.Delete().Where(category.ID(ID)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
