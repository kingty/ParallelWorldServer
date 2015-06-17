//用户相关的model

package models
import (
    
    "github.com/astaxie/beego/orm"
	"time"
)



type UserDetail struct{
	Id int64 `orm:"pk;auto"`
	User *User `orm:"rel(fk);column(user_id)"`
	Username string `orm:"size(100)"`
	Gender int `orm:"default(0)"`
	Url string `orm:"null;size(100)"`
	Motto string `orm:"size(100)"`
	Stone_count int `orm:"default(0)"`
	Get_count int `orm:"default(0)"`
	Get_time time.Time `orm:"auto_now_add;type(date)"`
}


func (u *UserDetail) TableName() string {
    return "user_detail"
}

func NewUserDetail() *UserDetail {
	obj := new(UserDetail)
	return obj
}


/*
*不会取出User的值
*/
func  GetUserDetailById(id int64) (UserDetail,error){
	o := getORM()
	userDetail  := UserDetail{Id:id}
	err:=o.Read(&userDetail)
	if err == orm.ErrNoRows {
    		Log("查询不到")
	} else if err == orm.ErrMissPK {
	    Log("找不到主键")
	} else if err!=nil{
	    Log(err)
	}else{
		Log(userDetail)
		
		
	}
	
	return userDetail,err
}

func AddUserDetail(u *UserDetail)(int64, error){
	o := getORM()
	id,err := o.Insert(u)
	if err!=nil{
		Log(err)
		return 0,err
	}
	return id,err
}





func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(NewUserDetail())
	
}
