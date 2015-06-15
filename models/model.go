package models
import(
	"fmt"
    "github.com/astaxie/beego/orm"
	"time"
	"encoding/json"
	"strings"
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
	return TypeTextField
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

var _ Fielder = new(JsonField)
