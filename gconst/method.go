package gconst

import (
	"reflect"
)

type Methods struct {
	Object any
}

// 获取配置
func (t Methods) GetOption(id int64) Option {
	val := reflect.ValueOf(t.Object).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Type().String() == "gconst.Option" {
			option := val.Field(i).Interface().(Option)
			if option.ID == id {
				return option
			}
		}
	}
	return Option{}
}

// 获取配置
func (t Methods) GetOptions() []Option {
	data := make([]Option, 0)
	val := reflect.ValueOf(t.Object).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Type().String() == "gconst.Option" {
			option := val.Field(i).Interface().(Option)
			data = append(data, option)
		}
	}
	return data
}

// 获取名称
func (t Methods) GetFromName() []IDName {
	data := make([]IDName, 0)
	val := reflect.ValueOf(t.Object).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Type().String() == "gconst.Option" {
			option := val.Field(i).Interface().(Option)
			data = append(data, IDName{option.ID, option.Name})
		}
	}
	return data
}

// 获取文本
func (t Methods) GetFromText() []IDName {
	data := make([]IDName, 0)
	val := reflect.ValueOf(t.Object).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Type().String() == "gconst.Option" {
			option := val.Field(i).Interface().(Option)
			data = append(data, IDName{option.ID, option.Text})
		}
	}
	return data
}

// 是否存在
func (t Methods) Exist(id int64) bool {
	val := reflect.ValueOf(t.Object).Elem()
	for i := 0; i < val.NumField(); i++ {
		if val.Field(i).Type().String() == "gconst.Option" {
			option := val.Field(i).Interface().(Option)
			if option.ID == id {
				return true
			}
		}
	}
	return false
}
