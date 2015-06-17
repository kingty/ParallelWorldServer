package service

import(

	. "ParallelWorldServer/models"
	
	. "ParallelWorldServer/lib"
)

type UserService struct{
	CommomService
}

func NewUserService() *UserService {
	obj := new(UserService)
	return obj
}

func(u *UserService) Login(email string,passWord string) (bool,error){
	user,err := GetUserByEmail(email)
	if err != nil{
		return false,err
	}
	pw := Strtomd5(user.Password)
	if pw == passWord{
		return true,nil
	}else{
		return false,NewError(WRONG_PASSWORD,"密码错误！")
	}
	
	
}