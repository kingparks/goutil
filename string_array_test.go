package goutil

import (
	"fmt"
	"testing"
)

func TestStrArr_Random(t *testing.T) {
	a := StrArr{"a", "b", "c", "c", "1"}
	fmt.Println(a.RemoveDuplicate())
}
