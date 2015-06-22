package controllers

import (
	"github.com/astaxie/beego"
	. "ParallelWorldServer/lib"
	"fmt"
	"strings"
	"time"
)

type MainController struct {
	beego.Controller
}



const(
	BaseUrl = ""//项目根url
	Version_1_0 = "1.0"//api版本号
	
	LoginUrl = "/user/login"
	RegisterUrl = "/user/register"
)

/*
*检验请求的合法性
*url签名+时间戳的机制
*/
func (this *MainController) CheakRequest(url string)(bool,error){
	userId := this.GetString("userId")//用户id
	sign := this.GetString("sign")//app客户端计算出来的签名
	timestamp,timestamperr := this.GetInt64("timestamp")//app客户端生成的时间戳
	if timestamperr==nil{
		localTimestamp := time.Now().Unix()//获取本地的时间戳
		
		dexTime := localTimestamp-timestamp//获取时间戳的差值
		
		
		if(dexTime>0 && dexTime/60<10){//时间戳的差值在10分钟内为正常操作
			bm, err := GetCACHE()
			if(err==nil){
					token := bm.Get(userId)	//取出服务端缓存的token值
					if token!=nil{
						//对 url&token=xxx&timestamp=xxx  进行签名验证，不让token在url中传输，防止token被截取
						
						timestampStr := fmt.Sprintf("%d", timestamp)  
						value := Strtomd5(url+"&token="+token.(string)+"&timestamp="+timestampStr)
						
						if strings.EqualFold(value,sign){
							
							return true,nil
						}else{
							return false,NewError(SIGNERROR,"签名验证失败！")
						}
					}else{
							return false,NewError(NOTOKEN,"token 不存在")
					}
			}else{
				return false,err
			}
		}else{
			return false,NewError(TIMESTAMPOUT,"timestamp参数缺失 或者 请求timestamp已超时!")
		}
	}else{
		return false,NewError(TIMESTAMPERROR,"timestamp参数错误!")
	}
}


