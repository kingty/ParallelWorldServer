package lib

import(
//	"errors"
)

const (
	WRONG_PASSWORD = iota
	NOUSER
	
	
)

type AppError struct{
	Num int
	Message string
}

func  NewError(num int,message string) error{
	
	err := AppError{Num:num,Message:message} 
	
	return &err
	
}


func(u *AppError) Error() string{
	return u.Message
}
