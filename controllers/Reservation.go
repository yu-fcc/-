package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ReservationController struct {
	beego.Controller
}

func (c *ReservationController) Get() {
	c.TplName = "reservation.html"
}

const (
	PNAME = "patientname"
)

// 前端通过后端请求地址/getallreservation来调用该方法，通过医生姓名查找Reservateinfo表中的信息传给前端
func (c *ReservationController) GetReservation() {
	var reservate []models.Reservateinfo
	doctorname := c.GetSession(DOCTORNAME)
	fmt.Println(doctorname)
	o := orm.NewOrm()
	o.QueryTable("reservateinfo").Filter("doctor_name", doctorname).All(&reservate)
	var respList []interface{}
	for _, reservate := range reservate {
		respList = append(respList, reservate.ReservateToRespDesc())
	}
	c.Data["json"] = respList
	c.ServeJSON()
}

type RloginRequest struct {
	Pname string `json:"pname"`
}

// 前端通过后端请求地址/selectreservation来调用该方法来实现模糊查找，需要前端传过来的查找信息，该方法中调用了models中reservationinfo.go里的SelectReservate方法
func (c *ReservationController) SelectReservation() {
	var rloginRequest RloginRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &rloginRequest)
	if err != nil {
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	var reservate []orm.Params
	reservate = models.SelectReservate(rloginRequest.Pname)
	fmt.Println(reservate)
	c.Data["json"] = reservate
	c.ServeJSON()
}

// 前端通过后端请求地址/findbyname/:id来调用该方法，通过获取前端的id值来查找Reservateinfo中的信息，然后保存姓名
func (c *ReservationController) FindByName() {
	id, err1 := c.GetInt(":id")
	if err1 != nil { // 有错误就返回数据：获取参数失败
		errespose1 := util.NewError(507)
		c.Data["json"] = errespose1
		c.ServeJSON()
		return
	}
	var reservate models.Reservateinfo
	o := orm.NewOrm()
	error := o.QueryTable("reservateinfo").Filter("id", id).One(&reservate)
	if error != nil {
		// 查询出错
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	c.SetSession(PNAME, reservate.Name)
	fmt.Println(reservate.Name)
	c.ServeJSON()
}
