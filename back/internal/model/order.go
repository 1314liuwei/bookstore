package model

type Order struct {
	OId     int64
	Amount  float64
	Status  string
	BookIds []int64
}
