package model

type Order struct {
	Data []int
}
type OrderBookInfo struct {
	BookId int `json:"book"`
}

type OrderRemove struct {
	OId int64
}

type OrderUpdate struct {
	OId    int64
	Status string
}

type OrderQuery struct {
	OId int64
}

type OrderQueryAll struct {
	Page int
}

type OrderQueryAllOutput struct {
	Price     float64 `json:"price"`
	BookImage string  `json:"bookImage"`
	BookName  string  `json:"bookName"`
	BookID    int     `json:"book_id"`
	CreatedAt string  `json:"created_at"`
	Ebook     string  `json:"ebook"`
	Id        int     `json:"id"`
	Oid       int64   `json:"oid"`
	Status    string  `json:"status"`
	UpdateAt  string  `json:"update_at"`
	UserOrder int     `json:"user_order"`
}
