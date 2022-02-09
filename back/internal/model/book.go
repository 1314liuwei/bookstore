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
	Name        string
	Author      string
	Description string
	Cover       string
	Ebook       string
	Price       float32
	Category    string
}
