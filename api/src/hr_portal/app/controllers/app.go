package controllers

import (
	"github.com/revel/revel"
	"hr_portal/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	firstname := c.Params.Form.Get("firstname")
	surname := c.Params.Form.Get("surname")
	password := c.Params.Form.Get("password")
	email := c.Params.Form.Get("email")
	phone := c.Params.Form.Get("phone")

	models.CreateUser(firstname, surname, password, email, phone)
	return c.Render()
}
