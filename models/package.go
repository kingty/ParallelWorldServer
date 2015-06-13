//卡包相关的model

package models
import (
    "fmt"
    "github.com/astaxie/beego/orm"
	"time"
)

type Card struct(
	Id uint64
	user_id uint64
	Url string
	Title string
	Content string
	Isprivate int
	City string
	Status int
	Create_time time.Time
)

type Package struct{
	Id uint64
	User_id uint64
	Url string
	Packagename string
	Isprivate int
	Status int
	Create_time time.Time
}

type PackageCard struct{
	Id uint64
	Card_id uint64
	Package_id uint64
	Category int
}

type Theme struct{
	Id uint64
	Url string
	Bgurl string
	Packagename string
	Status int
	Create_time time.Time
} 

type ThemeCard struct{
	Id uint64
	Card_id uint64
	Theme_id uint64
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