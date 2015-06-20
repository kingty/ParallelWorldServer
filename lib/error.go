package lib

import(
//	"errors"
)

//错误类型代码定义
const (
	COMMOM = iota
	WRONG_PASSWORD
	NOUSER
	CACHERROR
	NOTOKEN
	EMAILEXIST
	
	
)
//自定义自己的error
type AppError struct{
	Num int//错误代码
	Message string//错误消息
}

//初始化一个错误
func  NewError(num int,message string) error{
	err := AppError{Num:num,Message:message} 
	return &err
}


//实现error接口
func(u *AppError) Error() string{
	return u.Message
}
