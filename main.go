package main

import (
//	_ "ParallelWorld/routers"
//	"github.com/astaxie/beego"
	"fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
	"ParallelWorldServer/models"
	"time"
)


func init() {
    orm.RegisterDriver("mysql", orm.DR_MySQL)

    orm.RegisterDataBase("default", "mysql", "root:@/test?charset=utf8")
}

func main() {
//	beego.Run()
	o := orm.NewOrm()
    o.Using("test") // 默认使用 default，你可以指定为其他数据库

    user := new (models.User)

    user.Email="398036899@qq.com"
	user.Register_time = time.Now()
    

    
    fmt.Println(o.Insert(user))

}

