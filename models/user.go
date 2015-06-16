//用户相关的model

package models
import (
    
    "github.com/astaxie/beego/orm"
	"time"
	"fmt"
)

type User struct{
	Id int64 `orm:"pk;auto"`
	Email string `orm:"size(100)"`
	Password string `orm:"size(50)"`
	Status int `orm:"default(0)"`
	Register_time time.Time
}

//定义表名
func (u *User) TableName() string {
    return "user"
}


func NewUser() *User {
	obj := new(User)
	return obj
}




func  GetUserById(id int64) (User,error){
	o := getORM()
	user  := User{Id:id}
	err:=o.Read(&user)
	if err == orm.ErrNoRows {
    	fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
	    fmt.Println("找不到主键")
	} else if err!=nil{
	    fmt.Println(err)
	}else{
		fmt.Println(user)
		
		
	}
	
	return user,err
}

func AddUser(u *User)(int64, error){
	o := getORM()
	id,err := o.Insert(u)
	if err!=nil{
		fmt.Println(err)
	}
	return id,err
}

func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(NewUser())
	
}