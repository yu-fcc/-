package controllers

import "github.com/astaxie/beego"

type MainPageController struct {
	beego.Controller
}

func (c *MainPageController) Get() {
	c.TplName = "mainPage.html"
}
