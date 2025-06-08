package controllers

import "github.com/astaxie/beego"

type PatientSystemController struct {
	beego.Controller
}

func (c *PatientSystemController) Get() {
	c.TplName = "patientSystem.html"
}
