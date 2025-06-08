package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CheckCardController struct {
	beego.Controller
}

func (c *CheckCardController) Get() {
	c.TplName = "checkCard.html"
}

// 获取Checkinfo中的信息传给前端，前端通过后端请求地址/getcheckcard来调用
func (c *CheckCardController) GetCheckCard() {
	var check []models.Checkinfo
	o := orm.NewOrm()
	o.QueryTable("checkinfo").All(&check)
	var respList []interface{}
	for _, check := range check {
		respList = append(respList, check.CheckToRespDesc())
	}
	c.Data["json"] = respList
	c.ServeJSON()
}

func (c *CheckCardController) Choose() {
	id, err1 := c.GetInt(":id")
	if err1 != nil { // 有错误就返回数据：获取参数失败
		errespose1 := util.NewError(507)
		c.Data["json"] = errespose1
		c.ServeJSON()
		return
	}
	fmt.Println(id)
	var check models.Checkinfo
	o := orm.NewOrm()
	o.QueryTable("checkinfo").Filter("id", id).One(&check)
	var reservate models.Reservateinfo
	var doctor models.Doctorinfo
	patientname := c.GetSession(PNAME)
	doctorname := c.GetSession(DOCTORNAME)
	o.QueryTable("reservateinfo").Filter("name", patientname).One(&reservate)
	o.QueryTable("doctorinfo").Filter("uname", doctorname).One(&doctor)
	models.InsertData(reservate.Name, doctor.Uname, check.Name, check.Price, check.Adress)
	var response util.APIResponse
	var checkpatient models.Checkpatientinfo
	o.QueryTable("checkpatientinfo").Filter("pname", patientname).One(&checkpatient)
	fmt.Println(checkpatient.Id)
	if checkpatient.Id > 0 {
		response = util.JSONResponse(210, "添加成功")
	} else {
		response = util.JSONResponse(510, "添加失败")
	}
	c.Data["json"] = response
	c.ServeJSON()
}

type CloginRequest struct {
	Information string `json:"information"`
}

// 前端通过后端请求地址/selectcheckcard来调用该方法来实现模糊查找，需要前端传过来的查找信息，该方法中调用了models中check.go里的SelectCheck方法
func (c *CheckCardController) SelectCheckCard() {
	var cloginRequest CloginRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &cloginRequest)
	if err != nil {
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	var check []orm.Params
	check = models.SelectCheck(cloginRequest.Information)
	fmt.Println(check)
	c.Data["json"] = check
	c.ServeJSON()
}
