package controllers

import (

	. "ParallelWorldServer/service"
	. "ParallelWorldServer/dto"
	. "ParallelWorldServer/lib"

)

type UserController struct {
	MainController
}

var userInterface UserInterface  = new(UserService)

func (this *UserController) Login() {
	
	id,err := userInterface.Login("kingty@mofunsky.com","4652660")
	loginMessage := NewLoginMessage()
	
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
	this.Data["json"] = loginMessage
    this.ServeJson()
	
}

func(this *UserController)Register(){
	
	value := Strtomd5("ca9faf97c43d0a0f06cb9af8d51a1d7f"+ "kingty@mofunsky.com")
	hastoken,err := this.CheakToken(32,value)
	if err==nil{
		if hastoken{
			_,err2 := userInterface.Register("kingty@mofunsky.com","4652660")
			if(err2==nil){
				this.Data["json"] = "seccess"
			}else{
				this.Data["json"] = err2.Error()
			}
		}else{
			this.Data["json"] = "unkonw error!"
		}
	}else{
		this.Data["json"] =err.Error()
	}
	
    this.ServeJson()
}