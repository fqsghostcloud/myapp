package controllers

import (
	"flag"
	"os"

	"github.com/golang/glog"

	"myapp/models/db"

	"github.com/astaxie/beego"
)

type PageInfoController struct {
	beego.Controller
}

func (this *PageInfoController) Get() {
	flag.Parse()
	glog.Info("_____________________________________________________________")
	userAgent := this.Ctx.Request.UserAgent()
	ip := this.Ctx.Input.IP()
	host, err := os.Hostname()

	if err != nil {
		this.Ctx.WriteString("get host name error: " + err.Error() + "\n")
		glog.Error("get host name error: " + err.Error() + "\n")
	} else {
		this.Ctx.WriteString("host: " + host + "\n")
		glog.Info("host: " + host + "\n")
	}
	this.Ctx.WriteString("address: " + ip + "\n")
	this.Ctx.WriteString("UserAgent: " + userAgent + "\n")
	this.Ctx.WriteString("名称: " + os.Getenv("OEM") + "\n")
	this.Ctx.WriteString("版本: " + os.Getenv("VER") + "\n")
	glog.Info("address: " + ip + "\n" + "UserAgent: " + userAgent + "\n" + "名称: " + os.Getenv("OEM") + "\n" + "版本: " + os.Getenv("VER") + "\n")

	isLogin := this.GetSession("isLogin")
	if isLogin != nil && isLogin.(bool) == true {
		user := new(db.User)
		username := this.GetSession("username").(string)
		userInfo := user.GetIntro(username)
		this.Ctx.WriteString("username:" + username + "\n")
		this.Ctx.WriteString("intro:" + userInfo.Intro + "\n")
		glog.Info("username:" + username + "\n" + "intro:" + userInfo.Intro + "\n")
	} else {
		this.Ctx.WriteString("login faild!")
		glog.Info("login faild!")
	}
	glog.Info("_____________________________________________________________")
	glog.Flush()
}
