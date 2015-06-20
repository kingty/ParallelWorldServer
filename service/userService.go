package service

import(

	. "ParallelWorldServer/models"
	
	. "ParallelWorldServer/lib"
//	"github.com/astaxie/beego/cache"
	"strconv"
	"strings"
)

//user业务
type UserService struct{
	CommomService
}

func NewUserService() *UserService {
	obj := new(UserService)
	return obj
}

//登陆
func(u *UserService) Login(email string,passWord string) (int64,error){
	
	user,err := GetUserByEmail(email)
	if err != nil{
		return 0,err
	}
	pw := Strtomd5(passWord)
	if strings.EqualFold(pw, user.Password){
		//将生成的token 存入memcache缓存中
		bm, err := GetCACHE()
		 if(err==nil){
			key := strconv.FormatInt(user.Id,10)
			value,_ := u.GetUserTokenString(user)
		
			Log(key+":"+value)
			bm.Put( key,value,1000*60*60)
			
			return user.Id,nil
		}else{
			
			return 0,err
		}
		
	}else{
		return 0,NewError(WRONG_PASSWORD,"密码错误！")
	}
	

	
}


//注册
func (u *UserService) Register(email string,passWord string)(int64 ,error){
	
	user,err := GetUserByEmail(email)
	if err != nil{
		user := NewUser()
		user.Email = email
		user.Password = passWord
		
		id, err2 := AddUser(user)
		if err2 == nil{
			return id,nil
		}else{
			return 0,NewError(COMMOM,"插入错误")
		}
		
	}else{
		if(user.Id>0){
			return 0,NewError(EMAILEXIST,"邮箱已注册！")
		}
		return 0,NewError(COMMOM,"未知错误")
	}
	
}
