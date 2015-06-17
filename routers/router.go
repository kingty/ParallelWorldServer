package routers

import (
	"ParallelWorldServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/user/login", &controllers.UserController{},"*:Login")
	 
}
