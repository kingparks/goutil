package chrome

import (
	"fmt"
	"testing"
)

func TestGetResult(t *testing.T) {
	sResult, err := GetResult("http://xxx.xxx.xxx.xxx", `https://baidu.com`, "")
	fmt.Println(sResult, err)
}
