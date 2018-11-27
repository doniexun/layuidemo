package routers

import (
	"github.com/doniexun/layuidemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{}, "*:Index")
	beego.Router("/hello", &controllers.MainController{}, "*:Hello")
}
