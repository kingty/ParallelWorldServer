//卡包相关的model

package models
import (
    
    "github.com/astaxie/beego/orm"
	"time"
)


type Theme struct{
	Id int64 `orm:"pk;auto"`
	Url string `orm:"null;size(100)"`
	Bgurl string `orm:"null;size(100)"`
	Packagename string `orm:"size(100)"`
	Status int `orm:"default(0)"`
	Create_time time.Time
	Cards []*Card  `orm:"rel(m2m);rel_through(ParallelWorldServer/models.ThemeCard)"`
} 


func (u *Theme) TableName() string {
    return "theme"
}

func NewTheme() *Theme {
	obj := new(Theme)
	return obj
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(NewTheme())
	
}
