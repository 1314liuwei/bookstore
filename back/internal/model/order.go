package model

type Order struct {
	Id     int64
	Amount float64
	Status string
	BookId int64
}
