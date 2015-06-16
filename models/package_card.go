//卡包相关的model

package models
import (
    
    "github.com/astaxie/beego/orm"

)


type PackageCard struct{
	Id int64 `orm:"pk;auto"`
	Card *Card `orm:"rel(fk);column(card_id)"`
	Package *Package `orm:"rel(fk);column(package_id)"`
	Category int `orm:"default(0)"`
}


func (u *PackageCard) TableName() string {
    return "package_card"
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(PackageCard))
	
}
