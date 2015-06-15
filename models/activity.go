//用户活动相关的model

package models
import (
    "fmt"
    "github.com/astaxie/beego/orm"
	"time"
)

type Follow struct{
	Id uint64 `orm:"pk;auto"`
	Follower *User `orm:"rel(fk);column(follower_id)"`
	Followeder *User `orm:"rel(fk);column(followeder_id)"`
	Isfollowed int `orm:"default(0)"`
	Follow_time time.Time
}


type Message struct{
	Id uint64 `orm:"pk;auto"`
	Sender *User `orm:"rel(fk);column(sender_id)"`
	Receiver *User `orm:"rel(fk);column(receiver_id)"`
	Card *Card  `orm:"rel(fk);column(card_id)"`
	Content string `orm:"size(65535)"`
	Send_time time.Time
	Status int`orm:"default(0)"`
}


type StoneLog{
	Id uint64 `orm:"pk;auto"`
	User *User `orm:"rel(fk);column(user_id)"`
	Log string 
	Log_time time.Time
	Stone_count int 
	Status int `orm:"default(0)"`
}

//定义表名
func (u *Follow) TableName() string {
    return "follow"
}

func (u *Message) TableName() string {
    return "message"
}

func (u *StoneLog) TableName() string {
    return "stone_log"
}

func NewFollow() *Follow {
	obj := new(Follow)
	return obj
}
func NewMessage() *Message {
	obj := new(Message)
	return obj
}
func NewStoneLog() *StoneLog {
	obj := new(StoneLog)
	return obj
}