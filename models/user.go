//用户相关的model

package models
import (
    "fmt"
    "github.com/astaxie/beego/orm"
	"time"
)

type User struct{
	Id uint64
	Email string
	Password string
	Register_time time.Time
}

type UserDetail struct{
	Id uint64
	User_id uint64
	Username string
	Status int
	Gender int
	Url string
	Motto string
	Stone_count int
	Get_count int
	Get_time time.Time
}

//定义表名
func (u *User) TableName() string {
    return "user"
}
func (u *UserDetail) TableName() string {
    return "user_detail"
}
