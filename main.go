package main

import (
//	_ "ParallelWorld/routers"
//	"github.com/astaxie/beego"


	"fmt"
//    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
	m "ParallelWorldServer/models"
	"time"


)



func main() {
//	beego.Run()
//	o := orm.NewOrm()
//    err := o.Using("test") // 默认使用 default，你可以指定为其他数据库

	
    
	
	
//    user.Email="398036899@qq.com"
//	user.Register_time = time.Now()
    

//    row,err:=o.Insert(user)
	
//    fmt.Println(err,row)
	user,err := m.GetUserById(1)
	
	if err!=nil{
		fmt.Println("nil")
	}
	fmt.Println(user.Email)
	
	newUser := m.NewUser()
	newUser.Email = "kingty6@gmal.com"
	newUser.Register_time = time.Now()
	
	m.AddUser(newUser)

}

