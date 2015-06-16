//卡包相关的model

package models
import (
    
    "github.com/astaxie/beego/orm"

)

type ThemeCard struct{
	Id int64 `orm:"pk;auto"`
	Card *Card `orm:"rel(fk);column(card_id)"`
	Theme *Theme `orm:"rel(fk);column(theme_id)"`
}

func (u *ThemeCard) TableName() string {
    return "theme_card"
}


func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(new(ThemeCard))
	
}
