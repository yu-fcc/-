package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.TplName = "user.html"
}

/*
	name := c.GetString("username")
	password := c.GetString("password")
	if name == "admin" && password == "123456" {
		c.Ctx.WriteString("登陆成功")
	} else {
		c.Ctx.WriteString("登陆失败")
	}
}*/
