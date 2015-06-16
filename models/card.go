//卡包相关的model

package models
import (
    
    "github.com/astaxie/beego/orm"
	"time"
)

type Card struct{
	Id int64 `orm:"pk;auto"`
	User *User `orm:"rel(fk);column(user_id)"`
	Url string `orm:"null;size(100)"`
	Title string `orm:"size(100)"`
	Content string `orm:"size(65535)"`
	Isprivate int `orm:"default(0)"`
	City string `orm:"size(20)"`
	Status int `orm:"default(0)"`
	Create_time time.Time
}


func (u *Card) TableName() string {
    return "card"
}

func NewCard() *Card {
	obj := new(Card)
	return obj
}


func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(NewCard())
	
}
