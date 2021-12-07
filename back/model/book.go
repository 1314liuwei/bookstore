package model

import (
	"back/repository/ent"
	"back/repository/ent/book"
	"context"
	"time"
)

type Book struct {
	Name         string    `json:"name,omitempty"`
	Price        int       `json:"price,omitempty"`
	SurplusCatch int       `json:"surplus_catch,omitempty"`
	Author       string    `json:"author,omitempty"`
	Describe     string    `json:"describe,omitempty"`
	EBook        string    `json:"e_book,omitempty"`
	Cover        string    `json:"cover,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	Category     string    `json:"category"`
}

func CreateBook(ctx context.Context, client *ent.Client,
	name, author, describe, ebook, cover string,
	price, surplusCatch int,
	c *ent.Category) (Book, error) {
	b, err := client.Book.Create().
		SetName(name).
		SetAuthor(author).
		SetDescription(describe).
		SetEbook(ebook).
		SetCover(cover).
		SetCategory(c).Save(ctx)
	if err != nil {
		return Book{}, err
	}
	return Book{
		Name:      b.Name,
		Author:    b.Author,
		EBook:     b.Ebook,
		Cover:     b.Cover,
		Category:  c.Name,
		CreatedAt: b.CreatedAt}, nil
}

func QueryBook(ctx context.Context, client *ent.Client, id int) ([]Book, error) {
	all, err := client.Book.Query().Where(book.ID(id)).All(ctx)
	if err != nil {
		return []Book{}, err
	}
	var result []Book

	for _, b := range all {
		result = append(result, Book{
			Name:      b.Name,
			Author:    b.Author,
			Describe:  b.Description,
			EBook:     b.Ebook,
			Cover:     b.Cover,
			CreatedAt: b.CreatedAt,
			Price:     b.Price,
			//Category:  b.QueryCategory().Only(ctx),
		})
	}
	return result, nil
}

func UpdateBook(ctx context.Context, client *ent.Client, id int,
	name, author, describe, ebook, cover string,
	price, surplusCatch int,
	c *ent.Category) error {
	_, err := client.Book.Update().Where(book.ID(id)).
		SetName(name).
		SetAuthor(author).
		SetDescription(describe).
		SetEbook(ebook).
		SetCover(cover).
		SetCategory(c).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(ctx context.Context, client *ent.Client, id int) error {
	_, err := client.Book.Delete().Where(book.ID(id)).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
