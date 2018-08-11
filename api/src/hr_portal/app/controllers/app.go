package controllers

import (
	"github.com/revel/revel"
	"hr_portal/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	username := c.Params.Form.Get("username")
	password := c.Params.Form.Get("password")

	models.CreateUser(username, password)
	return c.Render()
}
