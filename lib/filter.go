package lib


import (
	. "ParallelWorldServer/dto"
	"fmt"
	"strings"
	"time"
	"strconv"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
)


/**
*过滤器
*过滤需要sign认证的请求。
*/
var CheckSignFilter = func(ctx *context.Context){
	userId := ctx.Input.Query("userId")//用户id
	sign := ctx.Input.Query("sign")//app客户端计算出来的签名
	timestampstr := ctx.Input.Query("timestamp")//app客户端生成的时间戳
	if(strings.EqualFold(userId,"")||strings.EqualFold(sign,"")||strings.EqualFold(timestampstr,"")){
		loginMessage := NewLoginMessage()
		loginMessage.Statu = 0
		ctx.WriteString("{\"status\":0,\"message\":\"参数缺失\"}")
	}else{
		
		timestamp,timestamperr := strconv.ParseInt(timestampstr,10,0)
		
		if timestamperr==nil{
			localTimestamp := time.Now().Unix()//获取本地的时间戳
			
			dexTime := localTimestamp-timestamp//获取时间戳的差值
			
			
			if(dexTime>0 && dexTime/60<2){//时间戳的差值在2分钟内为正常操作,超过2分钟后的操作视为非正常操作
				bm, err := GetCACHE()
				if(err==nil){
						token := bm.Get(userId)	//取出服务端缓存的token值
						if token!=nil{
							//对 url&token=xxx&timestamp=xxx  进行签名验证，不让token在url中传输，防止token被截取
							
							timestampStr := fmt.Sprintf("%d", timestamp)  
							signKey := ctx.Request.RequestURI+"&token="+token.(string)+"&timestamp="+timestampStr
							
							if !Verify(signKey,sign){
								ctx.WriteString("{\"status\":0,\"message\":\"签名验证失败!\"}")
							}
						}else{
							ctx.WriteString("{sta\"status\"tus:0,\"message\":\"token 不存在!\"}")
						}
				}else{
					ctx.WriteString("{\"status\":0,\"message\":\""+err.Error()+"\"}")
				}
			}else{
				ctx.WriteString("{\"status\":0,\"message\":\"timestamp超时!\"}")
			}
		}else{
			ctx.WriteString("{\"status\":0,\"message\":\"timestamp错误!\"}")
		}
	}
}


var AllowCors = cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "DELETE", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})