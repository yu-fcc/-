package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PatientListController struct {
	beego.Controller
}

func (c *PatientListController) Get() {
	c.TplName = "patientList.html"
}

// 获取Patientinfo中的信息传给前端，前端通过后端请求地址/getallpatient来调用
func (c *PatientListController) GetAllPatient() {
	var patient []models.Patientinfo
	o := orm.NewOrm()
	o.QueryTable("patientinfo").All(&patient)
	var respList []interface{}
	for _, patient := range patient {
		respList = append(respList, patient.PatientToRespDesc())
		//fmt.Println(user.Id, user.Uname)
	}
	c.Data["json"] = respList
	c.ServeJSON()
}

type PlloginRequest struct {
	Pname string `json:"pname"`
}

// 前端通过后端请求地址/selectpatient来调用该方法来实现模糊查找，需要前端传过来的查找信息，该方法中调用了models中patient.go里的SelectPatient方法
func (c *PatientListController) SelectPatient() {
	var plloginRequest PlloginRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &plloginRequest)
	if err != nil {
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	var patient []orm.Params
	patient = models.SelectPatient(plloginRequest.Pname)
	fmt.Println(patient)
	c.Data["json"] = patient
	c.ServeJSON()
}

// 前端通过后端请求地址/deletepatient/:id来调用该方法，通过获取前端的id值来删除Patientinfo表中相应的信息
func (c *PatientListController) DeletePatient() {
	id, err := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	if err != nil {            // 有错误就返回数据：获取参数失败
		errespose := util.NewError(507)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	o := orm.NewOrm() // 创建一个orm对象
	// 调用 orm 的 Delete 方法，&models.Userinfo{Id: id} 表示删除的是哪一个跟数据库相关的模型以及限制条件
	_, err = o.Delete(&models.Patientinfo{Id: id})
	var response util.APIResponse
	if err != nil { // 如果添加错误就返回学信息："删除数据失败"
		response = util.JSONResponse(508, "删除数据失败")
		c.ServeJSON()
		return
	}
	response = util.JSONResponse(202, "删除成功")
	c.Data["json"] = response
	c.ServeJSON()
}
