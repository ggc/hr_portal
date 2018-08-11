package controllers

import (
	"github.com/revel/revel"
	"github.com/ggc/hr_portal/api/hr_portalapp"
)
type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	sql := "SELECT username from users"
	rows, err := app.DB.Query(sql)

	return c.Render(rows)
}
