package model

type ShoppingCartChangeBook struct {
	ShoppingCartId int64
	BookIds        []int64
	Price          int64
}
