package models

import "github.com/astaxie/beego/orm"

type Checkinfo struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Adress string `json:"adress"`
	//Doctors *Doctorinfo `orm:"rel(fk)"`
}

func (check *Checkinfo) CheckToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":     check.Id,
		"name":   check.Name,
		"price":  check.Price,
		"adress": check.Adress,
		//"doctor_id": check.Doctors.Id,
	}
	return respInfo
}
func SelectCheck(name string) []orm.Params {
	o := orm.NewOrm()
	var check []orm.Params
	_, _ = o.Raw("select * from Checkinfo where name like ?", "%"+name+"%").Values(&check)
	return check
}
