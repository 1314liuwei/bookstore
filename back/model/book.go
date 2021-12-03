package model

import (
	"back/repository/ent"
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
		SetPrice(price).
		SetSurplusCatch(surplusCatch).
		SetAuthor(author).
		SetDescribe(describe).
		SetEbook(ebook).
		SetCover(cover).
		AddCategory(c).Save(ctx)
	if err != nil {
		return Book{}, nil
	}
	return Book{
		Name:         b.Name,
		Author:       b.Author,
		SurplusCatch: b.SurplusCatch,
		Describe:     b.Describe,
		EBook:        b.Ebook,
		Cover:        b.Cover,
		Category:     c.Name,
		CreatedAt:    b.CreatedAt}, nil
}

//func QueryCategory(ctx context.Context, client *ent.Client, ID int) (Book, error) {
//	client.Book.Query().Where()
//}
