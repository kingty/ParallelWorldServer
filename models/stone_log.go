//用户活动相关的model

package models
import (
    "github.com/astaxie/beego/orm"
	"time"
)

type StoneLog struct{
	Id int64 `orm:"pk;auto"`
	User *User `orm:"rel(fk);column(user_id)"`
	Log string 
	Log_time time.Time
	Stone_count int 
	Status int `orm:"default(0)"`
}

func (u *StoneLog) TableName() string {
    return "stone_log"
}

func NewStoneLog() *StoneLog {
	obj := new(StoneLog)
	return obj
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(NewStoneLog())
	
}