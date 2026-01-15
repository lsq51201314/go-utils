package gjwt

import (
	"fmt"
	"testing"
)

func TestGjwt(t *testing.T) {
	j := New("123456")
	s,e, _ := j.Generate("123456")
	fmt.Println(s,e)
	fmt.Println(j.Validate(s))
}
