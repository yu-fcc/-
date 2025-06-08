package controllers

import "github.com/astaxie/beego"

type DoctorSystemController struct {
	beego.Controller
}

func (c *DoctorSystemController) Get() {
	c.TplName = "doctorSystem.html"
}
