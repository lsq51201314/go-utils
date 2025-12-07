package gjwt

import (
	"fmt"
	"testing"
)

func TestGjwt(t *testing.T) {
	j := New("123456")
	s, _ := j.Generate(123456)
	fmt.Println(s)
	fmt.Println(j.Validate(s))
}
