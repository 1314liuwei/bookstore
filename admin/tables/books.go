package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetBooksTable(ctx *context.Context) table.Table {

	books := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := books.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Author", "author", db.Varchar)
	info.AddField("Description", "description", db.Longtext)
	info.AddField("Ebook", "ebook", db.Varchar)
	info.AddField("Cover", "cover", db.Varchar)
	info.AddField("Price", "price", db.Bigint)
	info.AddField("Created_at", "created_at", db.Timestamp)
	info.AddField("Category_book", "category_book", db.Bigint)

	info.SetTable("books").SetTitle("Books").SetDescription("Books")

	formList := books.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)
	formList.AddField("Author", "author", db.Varchar, form.Text)
	formList.AddField("Description", "description", db.Longtext, form.RichText)
	formList.AddField("Ebook", "ebook", db.Varchar, form.Text)
	formList.AddField("Cover", "cover", db.Varchar, form.Text)
	formList.AddField("Price", "price", db.Bigint, form.Number)
	formList.AddField("Created_at", "created_at", db.Timestamp, form.Datetime)
	formList.AddField("Category_book", "category_book", db.Bigint, form.Number)

	formList.SetTable("books").SetTitle("Books").SetDescription("Books")

	return books
}
