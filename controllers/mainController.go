package controllers

import (
	"github.com/astaxie/beego"
	. "ParallelWorldServer/lib"
//	"github.com/astaxie/beego/cache"
	"strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

//func (c *MainController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplNames = "index.tpl"
//}
func (this *MainController) CheakToken(userId int64,token string)(bool,error){
	bm, err := GetCACHE()
	if(err==nil){
			key := strconv.FormatInt(userId,10)
			
		
			s := bm.Get(key)	
			if s!=nil{
				if strings.EqualFold(s.(string),token){
					return true,nil
				}else{
					return false,NewError(NOTOKEN,"token 不存在")
				}
			}else{
					return false,NewError(NOTOKEN,"token 不存在")
			}
	}else{
		return false,err
	}
}


