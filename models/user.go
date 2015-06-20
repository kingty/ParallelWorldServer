//用户相关的model

package models
import (
    "errors"
	. "ParallelWorldServer/lib"
    "github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"time"
	
	
)
const(
	nomarl_state = iota
	delete_state 
)

type User struct{
	Id int64 `orm:"pk;auto"`
	Email string `orm:"size(100)" valid:"Email"`
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


func checkUser(u *User)(err error){
	valid := validation.Validation{}
	b, _ := valid.Valid(u)
	if !b {
		for _, err := range valid.Errors {
			Log(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func  GetUserById(id int64) (User,error){
	o := GetORM()
	user  := User{Id:id}
	err:=o.Read(&user)
	if err == orm.ErrNoRows {
    		Log("查询不到")
	} else if err == orm.ErrMissPK {
	    Log("找不到主键")
	} else if err!=nil{
	    Log(err.Error())
	}else{
		Log(user)
		
		
	}
	
	return user,err
}

func  GetUserByEmail(email string) (User,error){
	o := GetORM()
	user := User{Email: email}
	err := o.Read(&user, "Email")
	if err == orm.ErrNoRows {
    		Log("查询不到")
	} else if err == orm.ErrMissPK {
	    Log("找不到主键")
	} else if err!=nil{
	    Log(err.Error())
	}else{
		Log(user)
		
		
	}
	
	return user,err
}

func AddUser(u *User)(int64, error){
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := GetORM()
	u.Password = Strtomd5(u.Password)
	u.Register_time = time.Now()
	id,err := o.Insert(u)
	if err!=nil{
		Log(err)
		return 0,err
	}
	return id,err
}


/*
*非物理删除
*/
func DeleteUser(u *User)(int64, error) {
	o := GetORM()
	u.Status = delete_state
	num,err := o.Update(u, "status")
	if err!=nil{
		Log(err)
		return 0,err
	}
	return num, err
}




func init() {
    // 需要在init中注册定义的model
    orm.RegisterModel(NewUser())
	
}

