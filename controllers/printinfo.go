package controllers

import (
	"os"

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
}
