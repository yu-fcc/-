package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type ManagerLoginController struct {
	beego.Controller
}

func (c *ManagerLoginController) Get() {
	c.TplName = "managerLogin.html"
}

type MloginRequest struct {
	Mname string `json:"mname"`
	Mpwd  string `json:"mpwd"`
}

const (
	MLOGIN = "mlogin"
)

// 前端通过后端请求地址/domlogin来调用该方法，通过前端传过来的姓名查找Managerinfo表中相对应的密码来验证是否匹配，来实现登录
func (c *ManagerLoginController) Login() {
	var mloginRequest MloginRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &mloginRequest)
	if err != nil {
		errespose1 := util.NewError(504)
		c.Data["json"] = errespose1
		c.ServeJSON()
		return
	}
	var manager models.Managerinfo
	o := orm.NewOrm()
	o.QueryTable("managerinfo").Filter("mname", mloginRequest.Mname).One(&manager)
	// 验证密码是否匹配
	err = bcrypt.CompareHashAndPassword([]byte(manager.Mpwd), []byte(mloginRequest.Mpwd))
	if err != nil {
		errespose2 := util.NewError(506)
		c.Data["json"] = errespose2
		c.ServeJSON()
		return
	}
	var response util.APIResponse
	if manager.Id > 0 {
		redisCache := util.NewRedisCache()
		redisCache.Set(PLOGIN, manager.Mname)
		response = util.JSONResponse(200, "登录成功")
	} else {
		response = util.JSONResponse(502, "登录失败：用户名或者密码错误")
	}
	c.Data["json"] = response
	c.ServeJSON()
}
