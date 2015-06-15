//卡包相关的model

package models
import (
    
    "github.com/astaxie/beego/orm"
	"time"
)

type Card struct{
	Id uint64 `orm:"pk;auto"`
	User *User `orm:"rel(fk);column(user_id)"`
	Url string `orm:"null;size(100)"`
	Title string `orm:"size(100)"`
	Content string `orm:"size(65535)"`
	Isprivate int `orm:"default(0)"`
	City string `orm:"size(20)"`
	Status int `orm:"default(0)"`
	Create_time time.Time
}

type Package struct{
	Id uint64 `orm:"pk;auto"`
	User *User `orm:"rel(fk);column(user_id)"`
	Url string `orm:"null;size(100)"`
	Packagename string  `orm:"size(100)"`
	Isprivate int `orm:"default(0)"`
	Status int `orm:"default(0)"`
	Create_time time.Time
	Cards []*Card  `orm:"rel(m2m);rel_through(ParallelWorldServer/models.PackageCard)"`
}

type PackageCard struct{
	Id uint64 `orm:"pk;auto"`
	Card *Card `orm:"rel(fk);column(card_id)"`
	Package *Package `orm:"rel(fk);column(package_id)"`
	Category int `orm:"default(0)"`
}

type Theme struct{
	Id uint64 `orm:"pk;auto"`
	Url string `orm:"null;size(100)"`
	Bgurl string `orm:"null;size(100)"`
	Packagename string `orm:"size(100)"`
	Status int `orm:"default(0)"`
	Create_time time.Time
	Cards []*Card  `orm:"rel(m2m);rel_through(ParallelWorldServer/models.ThemeCard)"`
} 

type ThemeCard struct{
	Id uint64 `orm:"pk;auto"`
	Card *Card `orm:"rel(fk);column(card_id)"`
	Theme *Theme `orm:"rel(fk);column(theme_id)"`
}

//定义表名
func (u *Card) TableName() string {
    return "card"
}
func (u *Package) TableName() string {
    return "package"
}
func (u *PackageCard) TableName() string {
    return "package_card"
}
func (u *Theme) TableName() string {
    return "theme"
}
func (u *ThemeCard) TableName() string {
    return "theme_card"
}

func NewCard() *Card {
	obj := new(Card)
	return obj
}
func NewPackage() *Package {
	obj := new(Package)
	return obj
}
func NewTheme() *Theme {
	obj := new(Theme)
	return obj
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(NewCard(), NewPackage(),NewTheme(),new(PackageCard),new(ThemeCard))
	
}
