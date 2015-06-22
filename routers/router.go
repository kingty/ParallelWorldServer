package routers

import (
	"ParallelWorldServer/controllers"
	"github.com/astaxie/beego"
)



func init() {
    beego.Router("/", &controllers.UserController{},"*:Login")
	beego.Router(controllers.LoginUrl, &controllers.UserController{},"*:Login")
	beego.Router(controllers.RegisterUrl, &controllers.UserController{},"*:Register")
	 
}
