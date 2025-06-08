package controllers

import (
	"firstDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PersonalCenterController struct {
	beego.Controller
}

func IsPatientLoggedIn(c *beego.Controller) bool {
	// 获取 session 中的用户名
	patientname := c.GetSession(PATIENTNAME)
	fmt.Println(patientname)
	// 判断用户名是否存在
	if patientname == nil {
		return false
	}
	return true
}
func (c *PersonalCenterController) Get() {
	if !IsPatientLoggedIn(&c.Controller) {
		c.Redirect("/patientLogin", 302)
		return
	}
	c.TplName = "personalCenter.html"
}

// 前端通过后端请求地址/getone来调用该方法，通过患者姓名查找Patientinfo表中的信息传给前端
func (c *PersonalCenterController) GetOne() {
	patientname := c.GetSession(PATIENTNAME)
	var patient []models.Patientinfo
	o := orm.NewOrm()
	o.QueryTable("patientinfo").Filter("pname", patientname).One(&patient)
	var respList []interface{}
	for _, patient := range patient {
		respList = append(respList, patient.PatientToRespDesc())
	}
	c.Data["json"] = respList
	c.ServeJSON()
}
