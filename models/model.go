package models
import(
	"fmt"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)
// A json field.
type JsonField struct {
	Name string
	Data string
}

func (e *JsonField) String() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func (e *JsonField) FieldType() int {
	return orm.TypeTextField
}

func (e *JsonField) SetRaw(value interface{}) error {
	switch d := value.(type) {
	case string:
		return json.Unmarshal([]byte(d), e)
	default:
		return fmt.Errorf("<JsonField.SetRaw> unknown value `%v`", value)
	}
}

func (e *JsonField) RawValue() interface{} {
	return e.String()
}


func getORM() orm.Ormer{
	o := orm.NewOrm()
	return o
}


func init() {
    orm.RegisterDriver("mysql", orm.DR_MySQL)

    orm.RegisterDataBase("default", "mysql", "root:@/test?charset=utf8")
}
//var _ Fielder = new(JsonField)
