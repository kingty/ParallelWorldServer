package controllers

import (
	"github.com/astaxie/beego"
	. "ParallelWorldServer/service"
	. "ParallelWorldServer/dto"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Login() {
	var userInterface UserInterface  = new(UserService)
	isLogin,err := userInterface.Login("398036899@qq.com","123")
	loginMessage := NewLoginMessage()
	
	
	if err != nil{
		loginMessage.Statu = 0
		loginMessage.Message = err.Error()
	}else{
		if isLogin{
			loginMessage.Statu = 0
			loginMessage.Message = "未知错误"
		}else{
			loginMessage.Statu = 1
			loginMessage.Message = "登陆成功"
		}
		
		
	}
	this.Data["json"] = loginMessage
    this.ServeJson()
	
}
