package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetCategoriesTable(ctx *context.Context) table.Table {

	categories := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := categories.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Bigint).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)

	info.SetTable("categories").SetTitle("Categories").SetDescription("Categories")

	formList := categories.GetForm()
	formList.AddField("Id", "id", db.Bigint, form.Default)
	formList.AddField("Name", "name", db.Varchar, form.Text)

	formList.SetTable("categories").SetTitle("Categories").SetDescription("Categories")

	return categories
}
