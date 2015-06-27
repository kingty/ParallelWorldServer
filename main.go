package main

import (
	_ "ParallelWorldServer/routers"
//	. "ParallelWorldServer/controllers"
//	. "ParallelWorldServer/dto"
	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/context"
	


//	"fmt"
////    "github.com/astaxie/beego/orm"
//    _ "github.com/go-sql-driver/mysql"
//	m "ParallelWorldServer/models"
//	"time"
	. "ParallelWorldServer/lib"
//	"reflect"

)






func main() {
	beego.InsertFilter("/user/*",beego.BeforeRouter,CheckSignFilter)//需要sign验证的请求过滤器
	beego.InsertFilter("*", beego.BeforeRouter, AllowCors)//请求方法过滤
	beego.Run()
}

