package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type PatientRegisterController struct {
	beego.Controller
}

func (c *PatientRegisterController) Get() {
	c.TplName = "patientRegister.html"
}

// 前端通过后端请求地址/dopregister来调用该方法，获取前端的信息，将该信息插入到Patientinfo表中
func (c *PatientRegisterController) Register() {
	var patient models.Patientinfo
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &patient)
	if err != nil {
		errespose1 := util.NewError(504)
		c.Data["json"] = errespose1
		c.ServeJSON()
		return
	}
	// 生成密码哈希
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(patient.Ppwd), bcrypt.DefaultCost)
	if err != nil {
		errespose2 := util.NewError(505)
		c.Data["json"] = errespose2
		c.ServeJSON()
		return
	}
	patient.Ppwd = string(passwordHash)
	om := orm.NewOrm()
	id, err := om.Insert(&patient)
	var response util.APIResponse
	if id > 0 {
		//c.Data["json"] = LoginResponse{Success: true, Message: "注册成功"}
		response = util.JSONResponse(201, "注册成功")
	} else {
		//c.Data["json"] = LoginResponse{Success: false, Message: "注册失败"}
		response = util.JSONResponse(503, "注册失败")
	}
	c.Data["json"] = response
	c.ServeJSON()
}
