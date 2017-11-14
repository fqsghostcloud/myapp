package main

import (
	"myapp/models/grpc/service/server"
	_ "myapp/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.Session.SessionOn = true
	go beego.Run()
	server.Run()

	// beego.BConfig.WebConfig.Session.SessionOn = true
	// beego.Run()
}
