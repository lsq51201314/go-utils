package utils

import (
	"reflect"

	"github.com/spf13/cast"
)

// 常量配置
type ConstOptions struct {
	ID   int    `json:"id,string"`
	Name string `json:"name"`
	Text string `json:"text"`
}

// 常量实例
type Const struct{}

// 自动赋值
func (c Const) AutoBuild(dst ...interface{}) {
	for i := 0; i < len(dst); i++ {
		c.Default(dst[i])
	}
}

// 默认值
func (c Const) Default(t interface{}) {
	val := reflect.ValueOf(t).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		id := field.Tag.Get("id")
		val.Field(i).SetInt(cast.ToInt64(id))
	}
}

// 获取配置
func (c Const) GetOption(t interface{}, id int) (data ConstOptions) {
	val := reflect.ValueOf(t).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		if id == cast.ToInt(field.Tag.Get("id")) {
			data.ID = cast.ToInt(field.Tag.Get("id"))
			data.Name = field.Tag.Get("name")
			data.Text = field.Tag.Get("text")
			break
		}
	}
	return
}

// 获取配置
func (c Const) GetOptions(t interface{}) (data []ConstOptions) {
	data = make([]ConstOptions, 0)
	val := reflect.ValueOf(t).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		data = append(data, ConstOptions{
			ID:   cast.ToInt(field.Tag.Get("id")),
			Name: field.Tag.Get("name"),
			Text: field.Tag.Get("text"),
		})
	}
	return
}

// 获取配置
func (c Const) GetConfig(t interface{}, name string, id int) (config string) {
	val := reflect.ValueOf(t).Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		if id == cast.ToInt(field.Tag.Get("id")) {
			config = field.Tag.Get(name)
			return
		}
	}
	return ""
}
