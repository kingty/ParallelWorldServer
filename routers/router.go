package routers

import (
	"ParallelWorldServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.UserController{},"*:Login")
	beego.Router("/user/login", &controllers.UserController{},"*:Login")
	beego.Router("/user/register", &controllers.UserController{},"*:Register")
	 
}
