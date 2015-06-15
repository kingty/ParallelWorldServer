//用户相关的model

package models
import (
    
    "github.com/astaxie/beego/orm"
	"time"
)

type User struct{
	Id uint64 `orm:"pk;auto"`
	Email string `orm:"size(100)"`
	Password string `orm:"size(50)"`
	Status int `orm:"default(0)"`
	Register_time time.Time
}

type UserDetail struct{
	Id uint64 `orm:"pk;auto"`
	User *User `orm:"rel(fk);column(user_id)"`
	Username string `orm:"size(100)"`
	Gender int `orm:"default(0)"`
	Url string `orm:"null;size(100)"`
	Motto string `orm:"size(100)"`
	Stone_count int `orm:"default(0)"`
	Get_count int `orm:"default(0)"`
	Get_time time.Time `orm:"auto_now_add;type(date)"`
}

//定义表名
func (u *User) TableName() string {
    return "user"
}
func (u *UserDetail) TableName() string {
    return "user_detail"
}

func NewUser() *User {
	obj := new(User)
	return obj
}

func NewUserDetail() *UserDetail {
	obj := new(UserDetail)
	return obj
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(NewUser(), NewUserDetail())
	
}
