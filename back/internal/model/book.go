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
	Page     int
	ID       int
	Title    string
	Author   string
	Category string
}
