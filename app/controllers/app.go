package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	//return c.Render()
	greeting := "Aloha World"
	return c.Render(greeting)
}
