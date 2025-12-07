package gconst

import (
	"reflect"

	"github.com/spf13/cast"
)

type IDName struct {
	ID   int64  `json:"id,string"`
	Name string `json:"name"`
}

// 自动解析
func AutoUnmarshal(dst ...any) {
	for i := 0; i < len(dst); i++ {
		Unmarshal(dst[i])
	}
}

// 解析对象
func Unmarshal(t any) {
	val := reflect.ValueOf(t).Elem()
	for i := 0; i < val.NumField(); i++ {
		switch val.Field(i).Type().String() {
		case "gconst.Methods":
			_object := val.Field(i).FieldByName("Object")
			if _object.IsValid() && _object.CanSet() {
				val := reflect.ValueOf(t)
				_object.Set(val)
			}
		case "gconst.Option":
			field := val.Type().Field(i)

			_id := val.Field(i).FieldByName("ID")
			if _id.IsValid() && _id.CanSet() {
				val := field.Tag.Get("id")
				_id.SetInt(cast.ToInt64(val))
			}

			_type := val.Field(i).FieldByName("Type")
			if _type.IsValid() && _type.CanSet() {
				val := field.Tag.Get("type")
				_type.SetString(cast.ToString(val))
			}

			_name := val.Field(i).FieldByName("Name")
			if _name.IsValid() && _name.CanSet() {
				val := field.Tag.Get("name")
				_name.SetString(cast.ToString(val))
			}

			_text := val.Field(i).FieldByName("Text")
			if _text.IsValid() && _text.CanSet() {
				val := field.Tag.Get("text")
				_text.SetString(cast.ToString(val))
			}
		}
	}
}
