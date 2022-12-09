package go_util

import (
	"fmt"
	"testing"
)

func TestStr_GetBetween(t *testing.T) {
	fmt.Println(Str("1234").GetBetween("1", "232323"))
}
