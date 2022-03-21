package model

type BookInsert struct {
	Name        string
	Author      string
	Description string
	Cover       string
	Ebook       string
	Price       float32
}

type BookQueryInfo struct {
	ID       string
	Name     string
	Author   string
	Category string
}
