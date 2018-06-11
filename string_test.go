package util

import (
	"testing"
	"fmt"
)

func TestStr_GetBetween(t *testing.T) {
	fmt.Println(Str("1234").GetBetween("1","232323"))
}
