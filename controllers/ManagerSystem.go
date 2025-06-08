package controllers

import "github.com/astaxie/beego"

type ManagerSystemController struct {
	beego.Controller
}

func (c *ManagerSystemController) Get() {
	c.TplName = "managerSystem.html"
}
