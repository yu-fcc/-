package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type DoctorLoginController struct {
	beego.Controller
}

func (c *DoctorLoginController) Get() {
	c.TplName = "doctorLogin.html"
}

type DloginRequest struct {
	Uname string `json:"uname"`
	Upwd  string `json:"upwd"`
}

//type LoginResponse struct {
//	Success bool   `json:"success"`
//	Message string `json:"message"`
//}

const (
	DOCTORNAME = "doctorname"
	DLOGIN     = "dlogin"
)

// 前端通过后端请求地址/dodlogin来调用该方法，通过前端传过来的姓名查找Doctorinfo表中相对应的密码来验证是否匹配，来实现登录
func (c *DoctorLoginController) Login() {
	var dloginRequest DloginRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &dloginRequest)
	if err != nil {
		//c.Data["json"] = LoginResponse{Success: false, Message: "解析请求失败"}
		errespose1 := util.NewError(504)
		c.Data["json"] = errespose1
		c.ServeJSON()
		return
	}
	var doctor models.Doctorinfo
	o := orm.NewOrm()
	o.QueryTable("doctorinfo").Filter("uname", dloginRequest.Uname).One(&doctor)
	// 验证密码是否匹配
	err = bcrypt.CompareHashAndPassword([]byte(doctor.Upwd), []byte(dloginRequest.Upwd))
	if err != nil {
		//c.Data["json"] = map[string]string{"error": "Invalid username or password."}
		errespose2 := util.NewError(506)
		c.Data["json"] = errespose2
		c.ServeJSON()
		return
	}
	var response util.APIResponse
	//fmt.Println(doctor.Id)
	if doctor.Id > 0 {
		//保存用户名
		c.SetSession(DOCTORNAME, doctor.Uname)
		redisCache := util.NewRedisCache()
		redisCache.Set(DLOGIN, doctor.Uname)
		response = util.JSONResponse(200, "登录成功")
		//c.Data["json"] = LoginResponse{Success: true, Message: "登录成功"}
	} else {
		//c.Data["json"] = LoginResponse{Success: true, Message: "登录失败：用户名或者密码错误"}
		response = util.JSONResponse(502, "登录失败：用户名或者密码错误")
	}
	c.Data["json"] = response
	c.ServeJSON()
}
