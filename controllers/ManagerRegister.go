package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type ManagerRegisterController struct {
	beego.Controller
}

func (c *ManagerRegisterController) Get() {
	c.TplName = "managerRegister.html"
}

// 前端通过后端请求地址/domregister来调用该方法，获取前端的信息，将该信息插入到Managerinfo表中
func (c *ManagerRegisterController) Register() {
	var manager models.Managerinfo
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &manager)
	if err != nil {
		errespose1 := util.NewError(504)
		c.Data["json"] = errespose1
		c.ServeJSON()
		return
	}
	// 生成密码哈希
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(manager.Mpwd), bcrypt.DefaultCost)
	if err != nil {
		errespose2 := util.NewError(505)
		c.Data["json"] = errespose2
		c.ServeJSON()
		return
	}
	manager.Mpwd = string(passwordHash)
	om := orm.NewOrm()
	id, err := om.Insert(&manager)
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
