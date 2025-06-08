package controllers

import (
	"encoding/json"
	"firstDemo/models"
	"firstDemo/util"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
)

type DoctorListController struct {
	beego.Controller
}

func (c *DoctorListController) Get() {
	c.TplName = "doctorList.html"
}

// 获取Doctorinfo中的信息传给前端，前端通过后端请求地址/showlist来调用
func (c *DoctorListController) GetAll() {
	var doctor []models.Doctorinfo
	o := orm.NewOrm()
	o.QueryTable("doctorinfo").All(&doctor)
	var respList []interface{}
	for _, doctor := range doctor {
		respList = append(respList, doctor.DoctorToRespDesc())
		//fmt.Println(user.Id, user.Uname)
	}
	c.Data["json"] = respList
	c.ServeJSON()
}

//	func (c *DoctorListController) FindById() {
//		id, _ := c.GetInt(":id")
//		var doctor models.Doctorinfo
//		o := orm.NewOrm()
//		error := o.QueryTable("doctorinfo").Filter("id", id).One(&doctor)
//		if error != nil {
//			// 查询出错
//			errespose := util.NewError(504)
//			c.Data["json"] = errespose
//			c.ServeJSON()
//			return
//		}
//
//		// 将用户信息以JSON格式返回给前端
//		c.Data["json"] = doctor
//		c.ServeJSON()
//	}
//
// 前端通过后端请求地址/deleteDoctor/:id来调用该方法，通过获取前端的id值来删除Doctorinfo表中相应的信息
func (c *DoctorListController) DeleteDoctor() {
	id, err := c.GetInt(":id") // 获取到 url 当中 id 变量的值
	if err != nil {            // 有错误就返回数据：获取参数失败
		errespose1 := util.NewError(507)
		c.Data["json"] = errespose1
		c.ServeJSON()
		return
	}
	o := orm.NewOrm() // 创建一个orm对象
	// 调用 orm 的 Delete 方法，&models.Userinfo{Id: id} 表示删除的是哪一个跟数据库相关的模型以及限制条件
	_, err = o.Delete(&models.Doctorinfo{Id: id})
	var response1 util.APIResponse
	if err != nil { // 如果添加错误就返回学信息："删除数据失败"
		response1 = util.JSONResponse(508, "删除数据失败")
		c.ServeJSON()
		return
	}
	response1 = util.JSONResponse(202, "删除成功")
	c.Data["json"] = response1
	c.ServeJSON()
}

//func (c *DoctorListController) UpdateDoctor() {
//	var doctor models.Doctorinfo
//	err := json.Unmarshal(c.Ctx.Input.RequestBody, &doctor)
//	if err != nil {
//		errespose2 := util.NewError(504)
//		c.Data["json"] = errespose2
//		c.ServeJSON()
//		return
//	}
//	o := orm.NewOrm()
//	var r orm.RawSeter
//	r = o.Raw("UPDATE Doctorinfo SET  uname = ?,upwd=? WHERE id = ?", doctor.Uname, doctor.Upwd, doctor.Id)
//	_, error := r.Exec()
//	var response2 util.APIResponse
//	if error != nil {
//		response2 = util.JSONResponse(508, "更新失败")
//	} else {
//		response2 = util.JSONResponse(203, "更新成功")
//	}
//	c.Data["json"] = response2
//	c.ServeJSON()
//}

func (c *DoctorListController) ShowDoctorByPage() {
	var pageinfo models.PageParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pageinfo)
	if err != nil {
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	page := pageinfo.Pagenum
	pagesize := pageinfo.Pagesize
	doctors, err := models.GetUsers(page, pagesize)
	if err != nil {
		errespose := util.NewError(507)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	if len(doctors) > 0 {
		var repList []interface{}
		for _, doctor := range doctors {
			repList = append(repList, doctor.DoctorToRespDesc())
		}
		// 构建统一响应结构体
		response := models.UnifiedResponse{
			Code:    200,
			Message: "获取用户成功",
			Data:    repList,
		}
		c.Data["json"] = response
	}
	c.ServeJSON()
}
func (c *DoctorListController) GetDoctorList() {
	var pageinfo models.PageParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pageinfo)
	if err != nil {
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	currentPage := pageinfo.Pagenum
	pageSize := pageinfo.Pagesize
	fmt.Println(currentPage, pageSize)
	fmt.Println(pageinfo.Pagenum)
	o := orm.NewOrm()
	totalCount, _ := o.QueryTable("doctorinfo").Count()
	fmt.Println("总数量：", totalCount)
	totalPages := int(math.Ceil(float64(totalCount) / float64(pageSize)))
	fmt.Println("总数量：", totalPages)
	if currentPage < 1 {
		currentPage = 1
	} else if currentPage > totalPages {
		currentPage = totalPages
	}
	var doctors []models.Doctorinfo
	o.QueryTable("doctorinfo").Limit(pageSize, (currentPage-1)*pageSize).All(&doctors)
	var repList []interface{}
	for _, doctor := range doctors {
		repList = append(repList, doctor.DoctorToRespDesc())
	}
	//c.Data["json"] = repList
	// 构建分页结构体
	pagination := &models.Pagination{
		TotalCount:  totalCount,
		TotalPages:  totalPages,
		CurrentPage: currentPage,
		PageSize:    pageSize,
		Data:        repList,
	}

	// 返回分页信息给前端
	c.Data["json"] = pagination
	c.ServeJSON()
}
func (c *DoctorListController) ShowList() {
	var pageinfo models.PageParam
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &pageinfo)
	if err != nil {
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	currentPage := int64(pageinfo.Pagenum)
	pageSize := int64(pageinfo.Pagesize)
	fmt.Println(currentPage, pageSize)
	// 查询总记录数
	totalCount, _ := models.GetUserCount()
	// 创建分页实例
	pagination := util.NewPagination(totalCount, currentPage, pageSize)
	// 查询当前页的数据
	offset := (pagination.CurrentPage - 1) * pagination.PageSize
	doctors, _ := models.GetUserlist(offset, pagination.PageSize)
	var repList []interface{}
	for _, doctor := range doctors {
		repList = append(repList, doctor.DoctorToRespDesc())
	}
	pagination.Data = repList
	pagination.CurrentPage = currentPage
	fmt.Println(pagination.Data)
	fmt.Println(pagination.CurrentPage)
	fmt.Println(pagination.TotalPages)
	// 返回分页信息给前端
	c.Data["json"] = pagination
	c.ServeJSON()
}

type loginRequest struct {
	Dname string `json:"dname"`
}

// 前端通过后端请求地址/selectdoctor来调用该方法来实现模糊查找，需要前端传过来的查找信息，该方法中调用了models中doctor.go里的SelectDoctor方法
func (c *DoctorListController) SelectDoctor() {
	var dloginRequest loginRequest
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &dloginRequest)
	if err != nil {
		errespose := util.NewError(504)
		c.Data["json"] = errespose
		c.ServeJSON()
		return
	}
	var doctor []orm.Params
	doctor = models.SelectDoctor(dloginRequest.Dname)
	fmt.Println(doctor)
	c.Data["json"] = doctor
	c.ServeJSON()
}
