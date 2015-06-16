//用户活动相关的model

package models
import (
    "github.com/astaxie/beego/orm"
	"time"
)

type Follow struct{
	Id int64 `orm:"pk;auto"`
	Follower *User `orm:"rel(fk);column(follower_id)"`
	Followeder *User `orm:"rel(fk);column(followeder_id)"`
	Isfollowed int `orm:"default(0)"`
	Follow_time time.Time
}

func (u *Follow) TableName() string {
    return "follow"
}


func NewFollow() *Follow {
	obj := new(Follow)
	return obj
}
func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(NewFollow())
	
}