package gformat

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestGformat(t *testing.T) {
	fmt.Println(GetFloat64Null(0))
	fmt.Println(GetNullFloat64(sql.NullFloat64{}))

	fmt.Println(GetInt64Null(0))
	fmt.Println(GetNullInt64(sql.NullInt64{}))

	fmt.Println(GetInt32Null(0))
	fmt.Println(GetNullInt32(sql.NullInt32{}))

	fmt.Println(GetInt16Null(0))
	fmt.Println(GetNullInt16(sql.NullInt16{}))

	fmt.Println(GetStrNull(""))
	fmt.Println(GetNullStr(sql.NullString{}))

	fmt.Println(GetPinyinArr("你好啊"))
	fmt.Println(GetPinyin("你好啊"))
	fmt.Println(GetPinyinFirst("你好啊"))
}
