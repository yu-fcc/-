package models

type Managerinfo struct {
	Id    int    `json:"id"`
	Mname string `json:"mname"`
	Mpwd  string `json:"mpwd"`
}

func (manager *Managerinfo) ManagerToRespDesc() interface{} {
	respInfo := map[string]interface{}{
		"id":    manager.Id,
		"mname": manager.Mname,
	}
	return respInfo
}
