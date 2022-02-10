package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetOrdersTable(ctx *context.Context) table.Table {

	orders := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := orders.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Amount", "amount", db.Float)
	info.AddField("Status", "status", db.Enum)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Update_at", "update_at", db.Timestamp)
	info.AddField("Book_order", "book_order", db.Bigint)
	info.AddField("User_order", "user_order", db.Bigint)

	info.SetTable("orders").SetTitle("Orders").SetDescription("Orders")

	formList := orders.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Amount", "amount", db.Float, form.Text)
	formList.AddField("Status", "status", db.Enum, form.Text)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Update_at", "update_at", db.Timestamp, form.Datetime)
	formList.AddField("Book_order", "book_order", db.Bigint, form.Number)
	formList.AddField("User_order", "user_order", db.Bigint, form.Number)

	formList.SetTable("orders").SetTitle("Orders").SetDescription("Orders")

	return orders
}
