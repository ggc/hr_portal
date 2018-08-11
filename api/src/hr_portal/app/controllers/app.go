package controllers

import (
	"github.com/revel/revel"
	"hr_portal/app/models"

)
type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	models.CreateUser("Shine")
	return c.Render()
}
