package gconst

import (
	"errors"
	"strings"

	"github.com/spf13/cast"
)

type Option struct {
	ID   int64  `json:"id,string"`
	Type string `json:"type"`
	Name string `json:"name"`
	Text string `json:"text"`
}

// 获取名称
func (p Option) GetFromName() IDName {
	return IDName{p.ID, p.Name}
}

// 获取文本
func (p Option) GetFromText() IDName {
	return IDName{p.ID, p.Text}
}

// 格式化文本
func (p Option) FormatText(format ...string) string {
	text := p.Text
	if len(format) > 0 {
		for i := 0; i < len(format); i++ {
			text = strings.Replace(text, "%s", format[i], 1)
		}
	}
	return text
}

// 格式化错误
func (p Option) FormatError(format ...string) error {
	str := "(" + cast.ToString(p.ID) + ")" + p.FormatText(format...)
	return errors.New(str)
}
