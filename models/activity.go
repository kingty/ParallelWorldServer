//用户活动相关的model

package models
import (
    "fmt"
    "github.com/astaxie/beego/orm"
	"time"
)

type Follow struct{
	Id uint64
	Follower_id uint64
	Followeder_id uint64
	Isfollowed int
	Follow_time time.Time
}


type Message struct{
	Id uint64
	Sender_id uint64
	Receiver_id uint64
	Card_id uint64
	Content string
	Send_time time.Time
	Status int
}


type StoneLog{
	Id uint64
	User_id uint64
	Log string
	Log_time time.Time
	Stone_count int
	Status int
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