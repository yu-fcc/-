package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CheckResultController struct {
	beego.Controller
}

func (c *CheckResultController) Get() {
	c.TplName = "checkResult.html"
}

// 前端通过后端请求地址/getcheckresult来调用该方法，通过医生姓名查找Checkpatientinfo表中的信息传给前端
func (c *CheckResultController) GetCheckResult() {
	var checkpatient []models.Checkpatientinfo
	doctorname := c.GetSession(DOCTORNAME)
	o := orm.NewOrm()
	o.QueryTable("checkpatientinfo").Filter("dname", doctorname).All(&checkpatient)
	var respList []interface{}
	for _, checkpatient := range checkpatient {
		respList = append(respList, checkpatient.CheckPatientToRespDesc())
	}
	c.Data["json"] = respList
	c.ServeJSON()
}

type CrloginRequest struct {
	Information string `json:"information"`
}

// 前端通过后端请求地址/selectcheckresult来调用该方法来实现模糊查找，需要前端传过来的查找信息，该方法中调用了models中checkpatient.go里的SelectCheckResult方法
func (c *CheckResultController) SelectCheckResult() {
	var crloginRequest CrloginRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &crloginRequest)
	if err != nil {
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	var checkpatient []orm.Params
	checkpatient = models.SelectCheckResult(crloginRequest.Information)
	fmt.Println(checkpatient)
	c.Data["json"] = checkpatient
	c.ServeJSON()
}

// 前端通过后端请求地址/deletecheckresult/:id来调用该方法，通过获取前端的id值来删除Checkpatientinfo表中相应的信息
func (c *CheckResultController) DeleteCheckResult() {
	id, err := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	if err != nil {            // 有错误就返回数据：获取参数失败
		errespose := util.NewError(507)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	o := orm.NewOrm() // 创建一个orm对象
	// 调用 orm 的 Delete 方法，&models.Userinfo{Id: id} 表示删除的是哪一个跟数据库相关的模型以及限制条件
	_, err = o.Delete(&models.Checkpatientinfo{Id: id})
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

const (
	CHECKPNAME = "checkpname"
)

// 前端通过后端请求地址/getpname/:id来调用该方法，通过获取前端的id值来查找Checkpatientinfo表中相应的信息，再将该患者的姓名保存起来
func (c *CheckResultController) GetPname() {
	id, err := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	if err != nil {            // 有错误就返回数据：获取参数失败
		errespose := util.NewError(507)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	var checkpatient models.Checkpatientinfo
	o := orm.NewOrm()
	o.QueryTable("checkpatientinfo").Filter("id", id).One(&checkpatient)
	c.SetSession(CHECKPNAME, checkpatient.Pname)
}
