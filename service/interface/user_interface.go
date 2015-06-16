package models
import(
	
	"ParallelWorldServer/models"
)

type UserInterface interface{
	GetById(uint64) (User err)
	CreateUser(interface{}) error
	
}