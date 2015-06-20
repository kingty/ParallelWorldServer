package lib
import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"log"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/cache"
)

//create md5 string
func Strtomd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

//password hash function
func Pwdhash(str string) string {
	return Strtomd5(str)
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