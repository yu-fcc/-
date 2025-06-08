package controllers

import (
	"firstDemo/models"
	"firstDemo/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UpdateCheckResultController struct {
	beego.Controller
}

func (c *UpdateCheckResultController) Get() {
	c.TplName = "updateCheckResult.html"
}

// 前端通过后端请求地址/getupdatecheckcard来调用该方法，查找Checkinfo表中的信息传给前端
func (c *UpdateCheckResultController) GetUpdateCheckCard() {
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

// 前端通过后端请求地址/updatecheckresult/来调用该方法，通过前端获取的id查找Checkinfo表中信息，通过CheckResult保存的姓名来更新表中该姓名对应的信息
func (c *UpdateCheckResultController) UpdateCheckResult() {
	id, err1 := c.GetInt(":id")
	if err1 != nil { // 有错误就返回数据：获取参数失败
		errespose := util.NewError(507)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	fmt.Println(id)
	var check models.Checkinfo
	o := orm.NewOrm()
	o.QueryTable("checkinfo").Filter("id", id).One(&check)
	checkpname := c.GetSession(CHECKPNAME)
	var r orm.RawSeter
	r = o.Raw("UPDATE Checkpatientinfo SET  name = ?,price=?,adress=? WHERE pname = ?", check.Name, check.Price, check.Adress, checkpname)
	_, error := r.Exec()
	var response util.APIResponse
	if error != nil {
		response = util.JSONResponse(508, "更新失败")
	} else {
		response = util.JSONResponse(203, "更新成功")
	}
	c.Data["json"] = response
	c.ServeJSON()
}
