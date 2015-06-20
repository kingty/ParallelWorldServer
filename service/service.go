package service

import(
	
	. "ParallelWorldServer/models"
	. "ParallelWorldServer/lib"

)

//公共服务
type CommomService struct{
	
}

/**
*这里可以做一个权限检测，对每个Service.go中的业务进行封装一些公共操作
*/

func(this *CommomService) GetUserTokenString (u User)(string,error){
	if(u.Email==""||u.Password==""){
		return "",NewError(COMMOM,"密码和邮箱为空")
	}
	return Strtomd5(u.Password+ u.Email),nil
	
}