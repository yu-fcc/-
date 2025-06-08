package controllers

import (
	"firstDemo/models"
	"firstDemo/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ReservateInformationController struct {
	beego.Controller
}

func (c *ReservateInformationController) Get() {
	c.TplName = "reservateInformation.html"
}

// 前端通过后端请求地址/getreservate来调用该方法，通过患者姓名查找Reservateinfo表中的信息传给前端
func (c *ReservateInformationController) GetOneReservate() {
	patientname := c.GetSession(PATIENTNAME)
	fmt.Println(patientname)
	var reservate []models.Reservateinfo
	o := orm.NewOrm()
	o.QueryTable("reservateinfo").Filter("name", patientname).All(&reservate)
	var respList []interface{}
	for _, reservate := range reservate {
		respList = append(respList, reservate.ReservateToRespDesc())
	}
	c.Data["json"] = respList
	c.ServeJSON()
}

// 前端通过后端请求地址/deletereservate/:id来调用该方法，通过获取前端的id值来删除Reservateinfo表中相应的信息
func (c *ReservateInformationController) DeleteReservate() {
	id, err := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	if err != nil {            // 有错误就返回数据：获取参数失败
		errespose := util.NewError(507)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	o := orm.NewOrm() // 创建一个orm对象
	// 调用 orm 的 Delete 方法，&models.Userinfo{Id: id} 表示删除的是哪一个跟数据库相关的模型以及限制条件
	_, err = o.Delete(&models.Reservateinfo{Id: id})
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
