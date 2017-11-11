package controllers

import (
	"os"

	"myapp/models/db"

	"github.com/astaxie/beego"
)

type PageInfoController struct {
	beego.Controller
}

func (this *PageInfoController) Get() {
	userAgent := this.Ctx.Request.UserAgent()
	ip := this.Ctx.Input.IP()
	host, err := os.Hostname()

	if err != nil {
		this.Ctx.WriteString("get host name error: " + err.Error() + "\n")
	} else {
		this.Ctx.WriteString("host: " + host + "\n")
	}
	this.Ctx.WriteString("address: " + ip + "\n")
	this.Ctx.WriteString("UserAgent: " + userAgent + "\n")
	this.Ctx.WriteString("OEM: " + os.Getenv("OEM") + "\n")
	this.Ctx.WriteString("VER: " + os.Getenv("VER") + "\n")

	isLogin := this.GetSession("isLogin")
	if isLogin != nil && isLogin.(bool) == true {
		user := new(db.User)
		username := this.GetSession("username").(string)
		userInfo := user.GetIntro(username)
		this.Ctx.WriteString("username:" + username + "\n")
		this.Ctx.WriteString("intro:" + userInfo.Intro + "\n")
	} else {
		this.Ctx.WriteString("login faild!")
	}
}
