package gfile

import (
	"fmt"
	"testing"
)

func TestGfile(t *testing.T) {
	files := GetFiles("./")
	fmt.Println(files)
	fmt.Println(CreateMutiDir("./a/b/c"))
	fmt.Println(IsExist("./a/b/c"))
}
