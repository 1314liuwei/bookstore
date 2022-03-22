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
	BookIds []int `v:"required" json:"books"`
}
type ShoppingCartAddBooksRes struct{}

type ShoppingCartRemoveBooksReq struct {
	g.Meta  `path:"/shopping_cart/removeBook" method:"POST" tags:"shopping_cart" summary:"Empty books"`
	BookIds []int `v:"required" json:"books"`
}
type ShoppingCartRemoveBooksRes struct{}

type ShoppingCartEmptyReq struct {
	g.Meta `path:"/shopping_cart/empty" method:"POST" tags:"shopping_cart" summary:"Empty an shopping cart"`
}
type ShoppingCartEmptyRes struct{}

type ShoppingCartSettleReq struct {
	g.Meta  `path:"/shopping_cart/empty" method:"POST" tags:"shopping_cart" summary:"Empty an shopping cart"`
	BookIds []int `v:"required" json:"books"`
}
type ShoppingCartSettleRes struct{}
