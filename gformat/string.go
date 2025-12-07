package gformat

import (
	"database/sql"

	"github.com/mozillazg/go-pinyin"
)

func GetStrNull(s string) sql.NullString {
	r := sql.NullString{
		String: "",
		Valid:  false,
	}
	if s != "" {
		r.String = s
		r.Valid = true
	}
	return r
}

func GetNullStr(s sql.NullString) string {
	if !s.Valid {
		return ""
	}
	return s.String
}

func GetPinyinArr(hz string) [][]string {
	p := pinyin.NewArgs()
	return pinyin.Pinyin(hz, p)
}

func GetPinyin(hz string) []string {
	a := GetPinyinArr(hz)
	s := []string{}
	for _, v := range a {
		t := ""
		for _, x := range v {
			t += x
		}
		s = append(s, t)
	}
	return s
}

func GetPinyinFirst(hz string) string {
	p := GetPinyin(hz)
	s := ""
	for _, v := range p {
		if len(v) > 0 {
			s += v[:1]
		}
	}
	return s
}
