package goutil

import (
	"fmt"
	"github.com/Jetereting/goutil/net"
	"testing"
)

func TestStr_GetBetween(t *testing.T) {
	fmt.Println(Str("1234").GetBetween("1", "232323"))
}

func TestStr_GEO(t *testing.T) {
	fmt.Println(net.GetGEO("129.154.205.7"))
}
