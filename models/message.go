//用户活动相关的model

package models
import (
    "github.com/astaxie/beego/orm"
	"time"
)



type Message struct{
	Id int64 `orm:"pk;auto"`
	Sender *User `orm:"rel(fk);column(sender_id)"`
	Receiver *User `orm:"rel(fk);column(receiver_id)"`
	Card *Card  `orm:"rel(fk);column(card_id)"`
	Content string `orm:"size(65535)"`
	Send_time time.Time
	Status int`orm:"default(0)"`
}

func (u *Message) TableName() string {
    return "message"
}
func NewMessage() *Message {
	obj := new(Message)
	return obj
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel( NewMessage())
	
}