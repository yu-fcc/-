package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type PatientLoginController struct {
	beego.Controller
}

func (c *PatientLoginController) Get() {
	c.TplName = "patientLogin.html"
}

type PloginRequest struct {
	Pname string `json:"pname"`
	Ppwd  string `json:"ppwd"`
}

//type LoginResponse struct {
//	Success bool   `json:"success"`
//	Message string `json:"message"`
//}

const (
	PATIENTNAME = "patientname"
	PLOGIN      = "plogin"
)

// 前端通过后端请求地址/doplogin来调用该方法，通过前端传过来的姓名查找Patientinfo表中相对应的密码来验证是否匹配，来实现登录
func (c *PatientLoginController) Login() {
	var ploginRequest PloginRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ploginRequest)
	if err != nil {
		//c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	var patient models.Patientinfo
	o := orm.NewOrm()
	o.QueryTable("patientinfo").Filter("pname", ploginRequest.Pname).One(&patient)
	// 验证密码是否匹配
	err = bcrypt.CompareHashAndPassword([]byte(patient.Ppwd), []byte(ploginRequest.Ppwd))
	if err != nil {
		//c.Data["json"] = map[string]string{"error": "Invalid username or password."}
		errespose := util.NewError(506)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	var response util.APIResponse
	if patient.Id > 0 {
		c.SetSession(PATIENTNAME, patient.Pname)
		redisCache := util.NewRedisCache()
		redisCache.Set(PLOGIN, patient.Pname)
		response = util.JSONResponse(200, "登录成功")
		//errespose := util.NewError(200)
		//c.Data["json"] = errespose
	} else {
		//c.Data["json"] = LoginResponse{Success: true, Message: "登录失败：用户名或者密码错误"}
		response = util.JSONResponse(502, "登录失败：用户名或者密码错误")
		//errespose := util.NewError(502)
		//c.Data["json"] = errespose
	}
	c.Data["json"] = response
	c.ServeJSON()
}

//func (c *PatientLoginController) GetId() (id int) {
//	var ploginRequest PloginRequest
//	err := json.Unmarshal(c.Ctx.Input.RequestBody, &ploginRequest)
//	if err != nil {
//		//c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
//		errespose := util.NewError(504)
//		c.Data["json"] = errespose
//		c.ServeJSON()
//		return
//	}
//	var patient models.Patientinfo
//	o := orm.NewOrm()
//	o.QueryTable("patientinfo").Filter("pname", ploginRequest.Pname).One(&patient)
//	id = patient.Id
//	return id
//}
