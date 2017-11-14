package controllers

import (
	"myapp/models/db"
	"os"

	"github.com/astaxie/beego"
)

type SignUpController struct {
	beego.Controller
}

type User struct {
	Name  string `form:"username,text,用户名:"`
	Pwd   string `form:"pwd,password,密码:"`
	Intro string `form:"intro,textarea,个人介绍:"`
}

func (this *SignUpController) Get() {
	this.Data["Form"] = &User{}
	this.Data["AppInfo"] = os.Getenv("GOROOT")
	this.Layout = "layout.tpl"
	this.TplName = "signup.tpl"
	this.LayoutSections = make(map[string]string)
	this.Data["Oem"] = os.Getenv("OEM")
	this.Data["Ver"] = os.Getenv("VER")
	this.Layout = "layout.tpl"
	this.TplName = "signup.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["AppInfo"] = "appinfo.tpl"

}

func (this *SignUpController) Post() {
	status := false
	username := this.GetString("username")
	pwd := this.GetString("pwd")
	intro := this.GetString("intro")
	user := new(db.User)
	_, err := user.AddUser(username, pwd, intro)
	if err == nil {
		status = true
	}
	this.Data["Status"] = status
	this.Data["Oem"] = os.Getenv("OEM")
	this.Data["Ver"] = os.Getenv("VER")

	this.Layout = "layout.tpl"
	this.TplName = "notify.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["AppInfo"] = "appinfo.tpl"
}
