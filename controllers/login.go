package controllers

import (
	"myapp/models/db"

	"os"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

type LoginForm struct {
	Name string `form:"username,text,用户名:"`
	Pwd  string `form:"pwd,password,密码:"`
}

func (this *LoginController) Get() {
	this.Data["Form"] = &LoginForm{}
	this.Data["Oem"] = os.Getenv("OEM")
	this.Data["Ver"] = os.Getenv("VER")

	this.Layout = "layout.tpl"
	this.TplName = "login.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["AppInfo"] = "appinfo.tpl"
}

func (this *LoginController) Post() {
	username := this.GetString("username")
	pwd := this.GetString("pwd")
	user := new(db.User)
	status := user.CheckUser(username, pwd)
	if status {
		this.SetSession("username", username)
		this.SetSession("isLogin", true)
		this.Redirect("/user/profile/", 302)
	}
	this.Data["Status"] = status
	this.Data["Oem"] = os.Getenv("OEM")
	this.Data["Ver"] = os.Getenv("VER")

	this.Layout = "layout.tpl"
	this.TplName = "notify.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["AppInfo"] = "appinfo.tpl"
}
