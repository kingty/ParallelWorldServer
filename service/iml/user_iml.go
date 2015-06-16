package models

import(
	"fmt"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
	"ParallelWorldServer/models"
)
var o = orm.NewOrm()


func(u *models.User) GetById(id uint64)(models.User, error){	
o.Using("test")
	user := models.User{Id: id}
	err := o.Read(&user)

	if err == orm.ErrNoRows {
	    fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
	    fmt.Println("找不到主键")
	} else {
	    fmt.Println("GetById_Id=",user.Id)
	}
	return user,err

}