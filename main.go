package main

import (
	"firstDemo/models"
	_ "firstDemo/routers"
	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	//totalCount, _ := models.GetUserCount()
	//util.NewPagination(totalCount, 2, 10)
	//models.GetUserlist(10, 10)
	//models.GetUsers(2, 10)
	beego.Run()
}
