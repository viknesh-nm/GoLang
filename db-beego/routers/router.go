package routers

import (
	"github.com/astaxie/beego"
	"db-beego/controllers"
)

func init() {
	beego.Router("/", &controllers.ManageController{}, "*:Home")
	beego.Router("/add", &controllers.ManageController{}, "get,post:Add")
	beego.Router("/view", &controllers.ManageController{}, "get:View")
	beego.Router("/view/:id([0-9]+", &controllers.ManageController{}, "get:Vieww")
	beego.Router("/delete/:id([0-9]+)", &controllers.ManageController{}, "*:Delete")
	beego.Router("/update/:id([0-9]+)", &controllers.ManageController{}, "get,post:Update")
}
