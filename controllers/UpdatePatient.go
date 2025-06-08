package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpdatePatientController struct {
	beego.Controller
}

func (c *UpdatePatientController) Get() {
	id, _ := c.GetInt(":id")
	patient, _ := models.GetPatientById(id)
	fmt.Println(patient)
	if patient != nil {
		c.Data["patient"] = patient
	}
	fmt.Println(id)
	c.TplName = "updatePatient.html"
}

// 前端通过后端请求地址/doupdatepatient来调用该方法，通过从前端获取的信息来更新表中相应的信息
func (c *UpdatePatientController) UpdatePatient() {
	var patient models.Patientinfo
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &patient)
	if err != nil {
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("UPDATE Patientinfo SET  pname = ?,ppnb=?,padress=?,psex=?,pbirth=? WHERE id = ?", patient.Pname, patient.Ppnb, patient.Padress, patient.Psex, patient.Pbirth, patient.Id)
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
