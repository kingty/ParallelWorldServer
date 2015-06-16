//卡包相关的model

package models
import (
    
    "github.com/astaxie/beego/orm"
	"time"
)


type Package struct{
	Id int64 `orm:"pk;auto"`
	User *User `orm:"rel(fk);column(user_id)"`
	Url string `orm:"null;size(100)"`
	Packagename string  `orm:"size(100)"`
	Isprivate int `orm:"default(0)"`
	Status int `orm:"default(0)"`
	Create_time time.Time
	Cards []*Card  `orm:"rel(m2m);rel_through(ParallelWorldServer/models.PackageCard)"`
}

func (u *Package) TableName() string {
    return "package"
}
func NewPackage() *Package {
	obj := new(Package)
	return obj
}
func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(NewPackage())
	
}
