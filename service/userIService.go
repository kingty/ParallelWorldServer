package service
import(
	

)

//user业务层接口定义
type UserInterface interface{
	Login(email string,passWord string) (int64,error)
	Register(email string,passWord string)(int64 ,error)
	
	
}