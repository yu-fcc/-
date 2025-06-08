package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpdateReservateController struct {
	beego.Controller
}

func (c *UpdateReservateController) Get() {
	id, _ := c.GetInt(":id")
	//	fmt.Println("***********", id)
	reservate, _ := models.GetReservateById(id)
	//fmt.Println("**********", user)
	if reservate != nil {
		c.Data["reservate"] = reservate

	}
	c.TplName = "updateReservate.html"
}

// 前端通过后端请求地址/updatereservate来调用该方法，通过从前端获取的信息来更新表中相应的信息
func (c *UpdateReservateController) UpdateReservate() {
	var reservate models.Reservateinfo
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &reservate)
	if err != nil {
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return

	}
	o := orm.NewOrm()
	var r orm.RawSeter
	r = o.Raw("UPDATE reservateinfo SET  name = ?,time=?,doctor_name=? WHERE id = ?", reservate.Name, reservate.Time, reservate.DoctorName, reservate.Id)
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
