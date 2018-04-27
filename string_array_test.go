package util

import (
	"testing"
	"fmt"
)

func TestStrArr_Random(t *testing.T) {
	a:=StrArr{"a","b","c"}
	fmt.Println(a.Random2String())
}
