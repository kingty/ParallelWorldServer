package controllers

import (

	. "ParallelWorldServer/service"
	. "ParallelWorldServer/dto"
//	. "ParallelWorldServer/lib"
	"strings"
	
)

type UserController struct {
	MainController
}

var userInterface UserInterface  = new(UserService)

func (this *UserController) Login() {
	version := this.GetString("v")
//	device := this.GetString("device")
	
	loginMessage := NewLoginMessage()
	
	if(strings.EqualFold(Version_1_0,version)){
		id,err := userInterface.Login("kingty@mofunsky.com","xxx")
		if err != nil{
			loginMessage.Statu = 0
			loginMessage.Message = err.Error()
		}else{
			if id==0{
				loginMessage.Statu = 0
				loginMessage.Message = "未知错误"
			}else{
				loginMessage.Statu = 1
				loginMessage.Message = "登陆成功"
			}
			
			
		}
	}else{
		loginMessage.Statu = 0
		loginMessage.Message = "版本号错误！"
	}
	
	
	this.Data["json"] = loginMessage
    this.ServeJson()
	
}

func(this *UserController)Register(){
	version := this.GetString("v")
	if(strings.EqualFold(Version_1_0,version)){
		
		_,err2 := userInterface.Register("kingty4652660@gmail.com","11111")
		if(err2==nil){
			this.Data["json"] = "seccess"
		}else{
			this.Data["json"] = err2.Error()
		}
			
	}else{
		this.Data["json"] = "版本号有误"
	}
	
    this.ServeJson()
}
