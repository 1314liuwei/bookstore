package service

import (
	"back/internal/model"
	"context"
	"os"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sEBook struct {
}

var insEbook = sEBook{}

func EBook() *sEBook {
	return &insEbook
}

func (b sEBook) GetFile(ctx context.Context, in model.EBookGetFile) (string, error) {
	//uid := Context().Get(ctx).User.Id
	//all, err := dao.Orders.Ctx(ctx).Fields("book_order").Where(do.Orders{
	//	UserOrder: uid,
	//}).All()
	//if err != nil {
	//	return "", err
	//}

	flag := true
	//for _, record := range all {
	//	one, err := dao.Books.Ctx(ctx).Fields("ebook").Where(do.Books{
	//		Id: record["book_order"],
	//	}).One()
	//	if err != nil {
	//		return "", err
	//	}
	//
	//	if gconv.String(one["ebook"]) == in.File {
	//		flag = true
	//		break
	//	}
	//}

	if flag {
		pwd, _ := os.Getwd()
		return pwd + "/ebooks/" + in.File, nil
	} else {
		return "", gerror.New("you don't buy this ebook")
	}
}
