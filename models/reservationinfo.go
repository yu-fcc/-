package models

import "github.com/astaxie/beego/orm"

type Reservateinfo struct {
	Id         int    `json:"id"`
	Time       string `json:"time"`
	Name       string `json:"name"`
	DoctorName string `json:"dname"`
}

func (reservate *Reservateinfo) ReservateToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":    reservate.Id,
		"time":  reservate.Time,
		"name":  reservate.Name,
		"dname": reservate.DoctorName,
	}
	return respInfo
}
func GetReservateById(id int) (*Reservateinfo, error) {
	o := orm.NewOrm()
	reservater := Reservateinfo{Id: id}
	err := o.Read(&reservater)
	if err != nil {
		return nil, err
	}
	return &reservater, nil

}
func SelectReservate(name string) []orm.Params {
	o := orm.NewOrm()
	var reservate []orm.Params
	_, _ = o.Raw("select * from Reservateinfo where name like ?", "%"+name+"%").Values(&reservate)
	return reservate
}
