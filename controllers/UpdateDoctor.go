package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpdateDoctorController struct {
	beego.Controller
}

func (c *UpdateDoctorController) Get() {
	id, _ := c.GetInt(":id")
	doctor, _ := models.GetDoctorById(id)
	if doctor != nil {
		c.Data["doctor"] = doctor
	}
	c.TplName = "updateDoctor.html"
}

// 前端通过后端请求地址/doupdateusers来调用该方法，通过从前端获取的信息来更新表中相应的信息
func (c *UpdateDoctorController) UpdateUsers() {
	var doctor models.Doctorinfo
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &doctor)
	if err != nil {
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("UPDATE Doctorinfo SET  uname = ?,pnb=?,usex=? WHERE id = ?", doctor.Uname, doctor.Pnb, doctor.Usex, doctor.Id)
	_, error := r.Exec()
	var response util.APIResponse
	if error != nil {
		response = util.JSONResponse(508, "更新失败")
	} else {
		response = util.JSONResponse(203, "更新成功")
	}
	c.Data["json"] = response
	c.ServeJSON()
}
