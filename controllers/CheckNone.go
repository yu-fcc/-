package controllers

import (
	"firstDemo/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CheckNoneController struct {
	beego.Controller
}

func (c *CheckNoneController) Get() {
	c.TplName = "checkNone.html"
}

// 前端通过后端请求地址/getchecknone来调用该方法，通过患者姓名查找Checkpatientinfo表中的信息传给前端
func (c *CheckNoneController) GetCheckNone() {
	patientname := c.GetSession(PATIENTNAME)
	var checkpatient []models.Checkpatientinfo
	o := orm.NewOrm()
	o.QueryTable("checkpatientinfo").Filter("pname", patientname).One(&checkpatient)
	var respList []interface{}
	for _, checkpatient := range checkpatient {
		respList = append(respList, checkpatient.CheckPatientToRespDesc())
	}
	fmt.Println(respList)
	c.Data["json"] = respList
	c.ServeJSON()
}
