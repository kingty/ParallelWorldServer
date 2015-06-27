package service

import(
	
	. "ParallelWorldServer/models"
	. "ParallelWorldServer/lib"
	"strings"
)

//公共服务
type CommomService struct{
	
}

/**
*这里可以做一个权限检测，对每个Service.go中的业务进行封装一些公共操作
*/

func(this *CommomService) GetUserTokenString (u User)(string,error){
	if(strings.EqualFold(u.Email,"")){
		return "",NewError(COMMOM,"邮箱为空")
	}
	userDetail,err := GetUserDetailById(u.Id)
	if err!=nil{
		return "",err
	}
	
	return Pwdhash(u.Email+string(userDetail.Gender)+userDetail.Username+userDetail.Motto),nil
	
}