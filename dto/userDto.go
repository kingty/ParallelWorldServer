package dto

type LoginMessage struct{
	Statu int
	Message string
}
func NewLoginMessage() *LoginMessage {
	obj := new(LoginMessage)
	return obj
}