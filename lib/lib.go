package lib
import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"log"
	"strings"
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/cache"
	 "math/rand"
)



//生成随机加盐码加密的md5值
func GenerateMD5WithSalt(s string) (string){
	salt1 := fmt.Sprintf("%d", rand.Intn(99999999))
	salt2 := fmt.Sprintf("%d", rand.Intn(99999999))
	salt := salt1+salt2
	length := len(salt)
	if length<16{
		for i := 0; i < 16-length; i++ {
			salt = salt+"0"
		}
	}
	password := strtomd5(s+salt)
	
	passwordbyte := []byte(password)
	saltbyte := []byte(salt)
	var cs [48]byte
	for i := 0; i < 48; i += 3 {
            cs[i] = passwordbyte[i / 3 * 2];
            c := saltbyte[i / 3];
            cs[i + 1] = c;
            cs[i + 2] = passwordbyte[i / 3 * 2 + 1];
        } 
	return string(cs[:])
}

//验证密码正确与否
func Verify( password string,  md5Str string) bool{	
	cs1 := new ([32]byte)
	cs2 := new ([16]byte)
	md5byte :=  []byte(md5Str)
	for i := 0; i < 48; i += 3 {
       cs1[i / 3 * 2] = md5byte[i];
       cs1[i / 3 * 2 + 1] = md5byte[i + 2]
       cs2[i / 3] = md5byte[i + 1]
    } 
	return strings.EqualFold(strtomd5(password + string(cs2[:])),string(cs1[:]));
}

//create md5 string
func strtomd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}




//password hash function
func Pwdhash(str string) string {
	return strtomd5(str)
}

func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}



func GetORM() orm.Ormer{
	o := orm.NewOrm()
	return o
}



func Log(v ... interface{}){
	log.Println(v...)
}


func GetCACHE() (cache.Cache, error){
	bm, err := cache.NewCache("memory", `{"conn":"127.0.0.1:11211"}`)
	return bm,err
	
}