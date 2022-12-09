package go_util

import (
	"fmt"
	"testing"
)

func TestStrArr_Random(t *testing.T) {
	a := StrArr{"a", "b", "c"}
	fmt.Println(a.Random2String())
}
