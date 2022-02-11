package v1

import "github.com/gogf/gf/v2/frame/g"

type ShoppingCartCreateReq struct {
	g.Meta `path:"/shopping_cart/create" method:"POST" tags:"shopping_cart" summary:"Create an shopping cart"`
}
type ShoppingCartCreateRes struct {
	SId int64
}

type ShoppingCartAddBooksReq struct {
	g.Meta  `path:"/shopping_cart/addBook" method:"POST" tags:"shopping_cart" summary:"Add books"`
	SId     int64   `v:"required"`
	BookIds []int64 `v:"required"`
}
type ShoppingCartAddBooksRes struct{}

type ShoppingCartRemoveBooksReq struct {
	g.Meta  `path:"/shopping_cart/removeBook" method:"POST" tags:"shopping_cart" summary:"Remove books"`
	SId     int64   `v:"required"`
	BookIds []int64 `v:"required"`
}
type ShoppingCartRemoveBooksRes struct{}

type ShoppingCartRemoveReq struct {
	g.Meta `path:"/shopping_cart/remove" method:"POST" tags:"shopping_cart" summary:"Remove an shopping cart"`
	SId    int64 `v:"required"`
}
type ShoppingCartRemoveRes struct{}
