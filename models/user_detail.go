//用户相关的model

package models
import (
    
    "github.com/astaxie/beego/orm"
	"time"
)



type UserDetail struct{
	Id int64 `orm:"pk;auto"`
	User *User `orm:"rel(fk);column(user_id)"`
	Username string `orm:"size(100)"`
	Gender int `orm:"default(0)"`
	Url string `orm:"null;size(100)"`
	Motto string `orm:"size(100)"`
	Stone_count int `orm:"default(0)"`
	Get_count int `orm:"default(0)"`
	Get_time time.Time `orm:"auto_now_add;type(date)"`
}


func (u *UserDetail) TableName() string {
    return "user_detail"
}

func NewUserDetail() *UserDetail {
	obj := new(UserDetail)
	return obj
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(NewUserDetail())
	
}
