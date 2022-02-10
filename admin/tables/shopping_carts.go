package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetShoppingCartsTable(ctx *context.Context) table.Table {

	shoppingCarts := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := shoppingCarts.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Amount", "amount", db.Bigint)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Update_at", "update_at", db.Timestamp)
	info.AddField("Book_shopping_cart", "book_shopping_cart", db.Bigint)
	info.AddField("User_shopping_cart", "user_shopping_cart", db.Bigint)

	info.SetTable("shopping_carts").SetTitle("ShoppingCarts").SetDescription("ShoppingCarts")

	formList := shoppingCarts.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Amount", "amount", db.Bigint, form.Number)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Update_at", "update_at", db.Timestamp, form.Datetime)
	formList.AddField("Book_shopping_cart", "book_shopping_cart", db.Bigint, form.Number)
	formList.AddField("User_shopping_cart", "user_shopping_cart", db.Bigint, form.Number)

	formList.SetTable("shopping_carts").SetTitle("ShoppingCarts").SetDescription("ShoppingCarts")

	return shoppingCarts
}
