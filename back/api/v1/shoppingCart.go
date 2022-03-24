package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ShoppingCartQueryReq struct {
	g.Meta `path:"/shopping_cart/query" method:"GET" tags:"shopping_cart" summary:"Create an shopping cart"`
}
type ShoppingCartQueryRes struct {
	Result []int `json:"result"`
}

type ShoppingCartAddBooksReq struct {
	g.Meta  `path:"/shopping_cart/addBook" method:"POST" tags:"shopping_cart" summary:"Add books"`
	BookIds []int `v:"required"`
}
type ShoppingCartAddBooksRes struct{}

type ShoppingCartRemoveBooksReq struct {
	g.Meta  `path:"/shopping_cart/removeBook" method:"POST" tags:"shopping_cart" summary:"Empty books"`
	BookIds []int `v:"required"`
}
type ShoppingCartRemoveBooksRes struct{}

type ShoppingCartEmptyReq struct {
	g.Meta `path:"/shopping_cart/empty" method:"POST" tags:"shopping_cart" summary:"Empty an shopping cart"`
}
type ShoppingCartEmptyRes struct{}

type ShoppingCartSettleReq struct {
	g.Meta  `path:"/shopping_cart/empty" method:"POST" tags:"shopping_cart" summary:"Empty an shopping cart"`
	BookIds []int `v:"required"`
}
type ShoppingCartSettleRes struct{}
