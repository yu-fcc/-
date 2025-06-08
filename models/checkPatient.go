package models

import "github.com/astaxie/beego/orm"

type Checkpatientinfo struct {
	Id     int    `json:"id"`
	Pname  string `json:"pname"`
	Dname  string `json:"dname"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Adress string `json:"adress"`
}

func (checkpatient *Checkpatientinfo) CheckPatientToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":     checkpatient.Id,
		"pname":  checkpatient.Pname,
		"dname":  checkpatient.Dname,
		"name":   checkpatient.Name,
		"price":  checkpatient.Price,
		"adress": checkpatient.Adress,
	}
	return respInfo
}
func InsertData(pname string, dname string, name string, price int, adress string) {
	var checkpatient Checkpatientinfo
	o := orm.NewOrm()
	_ = o.Raw("insert into Checkpatientinfo(Pname,Dname,Name,Price,Adress) values(?,?,?,?,?) ", pname, dname, name, price, adress).QueryRow(&checkpatient)
}
func SelectCheckResult(name string) []orm.Params {
	o := orm.NewOrm()
	var checkpatient []orm.Params
	_, _ = o.Raw("select * from Checkpatientinfo where pname like ?", "%"+name+"%").Values(&checkpatient)
	return checkpatient
}
