package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"math"
)

type Doctorinfo struct {
	Id    int    `json:"id"`
	Uname string `json:"uname"`
	Upwd  string `json:"upwd"`
	Pnb   string `json:"pnb"`
	Usex  string `json:"usex"`
}

func (doctor *Doctorinfo) DoctorToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":    doctor.Id,
		"uname": doctor.Uname,
		"usex":  doctor.Usex,
		"pnb":   doctor.Pnb,
	}
	return respInfo
}

func SelectDoctor(name string) []orm.Params {
	o := orm.NewOrm()
	var doctor []orm.Params
	_, _ = o.Raw("select * from Doctorinfo where uname like ?", "%"+name+"%").Values(&doctor)
	return doctor
}

type PageParam struct {
	Pagesize int `json:"pagesize"` //每页显示多少条
	Pagenum  int `json:"pagenum"`  //第几页
}
type UnifiedResponse struct {
	Code    int         `json:"code"`    // 响应状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}
type Pagination struct {
	TotalCount  int64         `json:"total_count"`  // 总记录数
	TotalPages  int           `json:"total_pages"`  // 总页数
	CurrentPage int           `json:"current_page"` // 当前页数
	PageSize    int           `json:"page_size"`    // 每页记录数
	Data        []interface{} `json:"data"`         // 当前页的数据
}

func GetDoctorById(id int) (*Doctorinfo, error) {
	o := orm.NewOrm()
	doctor := Doctorinfo{Id: id}
	err := o.Read(&doctor)
	if err != nil {
		return nil, err
	}
	return &doctor, nil
}

func GetUsers(page, pageSize int) ([]Doctorinfo, error) {
	o := orm.NewOrm()
	count, _ := o.QueryTable("doctorinfo").Count()
	fmt.Println("总数量：", count)
	// 计算总页数
	totalPage := int(math.Ceil(float64(count) / float64(pageSize)))
	// 确保当前页码在有效范围内
	if page < 1 {
		page = 1
	} else if page > totalPage {
		page = totalPage
	}
	var doctors []Doctorinfo
	_, err := o.QueryTable("doctorinfo").Limit(pageSize, (page-1)*pageSize).All(&doctors)
	return doctors, err
}

// 获取总记录数
func GetUserCount() (int64, error) {
	o := orm.NewOrm()
	count, err := o.QueryTable("doctorinfo").Count()
	if err != nil {
		return 0, err

	}
	return count, nil
}

// 获取当前页的数据
func GetUserlist(offset, limit int64) ([]Doctorinfo, error) {
	o := orm.NewOrm()
	var doctors []Doctorinfo
	_, err := o.QueryTable("doctorinfo").Offset(offset).Limit(limit).All(&doctors)
	if err != nil {
		return nil, err

	}
	return doctors, nil
}
