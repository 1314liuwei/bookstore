package model

type Order struct {
	Data []OrderBookInfo
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
