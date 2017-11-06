package main

import (
	"myapp/models/grpc/service/server"
	_ "myapp/routers"

	"github.com/astaxie/beego"
)

func main() {
	go beego.Run()
	server.Run()
}
