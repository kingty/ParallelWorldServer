package main

import (
//	_ "ParallelWorld/routers"
//	"github.com/astaxie/beego"


	"fmt"
//    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
	m "ParallelWorldServer/models"
	"time"
	. "ParallelWorldServer/service"
	

)



func main() {
//	beego.Run()
//	o := orm.NewOrm()
//    err := o.Using("test") // 默认使用 default，你可以指定为其他数据库

	
    
	
	user := m.NewUser()
    user.Email="398036897qq.com"
	user.Password = "kingty"
	user.Register_time = time.Now()
    

    row,err := m.AddUser(user)
	
    fmt.Println(err,row)
	
	//接口
	var userInterface UserInterface  = new(UserService)
	
	islogin,err2 := userInterface.Login("398036899@qq.com","123")
	fmt.Println(islogin,err2)
	

	
//	userInterface.
	
//	user,err := m.GetUserById(1)
	
//	if err!=nil{
//		fmt.Println("nil")
//	}
//	fmt.Println(user.Email)
//	m.DeleteUser(&user)
	
	
//	newUser := m.NewUser()
//	newUser.Email = "kingty6@gmal.com"
//	newUser.Register_time = time.Now()
	
//	m.AddUser(newUser)
	
//	newUserDetail := m.NewUserDetail()
//	newUserDetail.Motto = "sssss"
//	newUserDetail.User=&user
//	newUserDetail.Username = "kingty"
	
//	m.AddUserDetail(newUserDetail)
	
//	userdetail,_ := m.GetUserDetailById(1)
////	fmt.Println(userdetail)
//	fmt.Println(userdetail.User.Email)

}

