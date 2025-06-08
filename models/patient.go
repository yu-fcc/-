package models

import "github.com/astaxie/beego/orm"

type Patientinfo struct {
	Id      int    `json:"id"`
	Pname   string `json:"pname"`
	Ppwd    string `json:"ppwd"`
	Ppnb    string `json:"ppnb"`
	Padress string `json:"padress"`
	Psex    string `json:"psex"`
	Pbirth  string `json:"pbirth"`
}

func (patient *Patientinfo) PatientToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":      patient.Id,
		"pname":   patient.Pname,
		"ppwd":    patient.Ppwd,
		"ppnb":    patient.Ppnb,
		"padress": patient.Padress,
		"psex":    patient.Psex,
		"pbirth":  patient.Pbirth,
	}
	return respInfo
}
func SelectPatient(name string) []orm.Params {
	o := orm.NewOrm()
	var patient []orm.Params
	_, _ = o.Raw("select * from Patientinfo where pname like ?", "%"+name+"%").Values(&patient)
	return patient
}
func GetPatientById(id int) (*Patientinfo, error) {
	o := orm.NewOrm()
	patient := Patientinfo{Id: id}
	err := o.Read(&patient)
	if err != nil {
		return nil, err
	}
	return &patient, nil

}
