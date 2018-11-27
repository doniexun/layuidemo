package main

import (
	_ "github.com/doniexun/layuidemo/routers"
	"github.com/doniexun/layuidemo/controllers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/layui", &controllers.MainController{}, "*:Index")
	beego.Run()
}

