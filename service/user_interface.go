package service
import(
	

)

type UserInterface interface{
	Login(email string,passWord string) (bool,error)
	
	
	
}