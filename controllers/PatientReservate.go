package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PatientReservateController struct {
	beego.Controller
}

func (c *PatientReservateController) Get() {
	c.TplName = "patientReservate.html"
}

// 前端通过后端请求地址/doreservate来调用该方法，获取前端的信息，将该信息插入到Reservateinfo表中
func (c *PatientReservateController) GetReservate() {
	var reservate models.Reservateinfo
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &reservate)
	if err != nil {
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	om := orm.NewOrm()
	id, err := om.Insert(&reservate)
	var response util.APIResponse
	if id > 0 {
		response = util.JSONResponse(204, "预约成功")
	} else {
		response = util.JSONResponse(504, "预约失败")
	}
	c.Data["json"] = response
	c.ServeJSON()
}
