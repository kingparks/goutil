package go_util

import (
	"fmt"
	"testing"
)

func TestT_ToInt(t *testing.T) {
	fmt.Println(NewT("32434").ToInt())
}

func TestT_ToString(t *testing.T) {
	fmt.Println(NewT("32434").ToString())
}

func TestT_ToSlice(t *testing.T) {
	fmt.Println(NewT("324,34").ToSlice())
}
