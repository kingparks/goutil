package censys

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/tidwall/gjson"
	"strings"
)

func Search(sURL, sName, sPwd string) (result []string, err error) {
	next := ""
run:
	resultS, err := httplib.
		Get(sURL+"&cursor="+next).
		SetBasicAuth(sName, sPwd).String()
	if err != nil {
		fmt.Println(err)
		return
	}
	gResult := gjson.Parse(resultS).Get("result")
	gResult.Get("hits").ForEach(func(key, value gjson.Result) bool {
		value.Get("services").ForEach(func(sKey, sValue gjson.Result) bool {
			resultC := ""
			// 协议
			proto := strings.ToLower(sValue.Get("extended_service_name").Str)
			switch proto {
			case "http":
				resultC += proto + "://"
			case "https":
				resultC += proto + "://"
			default:
				return true
			}
			// 域名或IP
			if value.Get("name").Str == "" {
				resultC += value.Get("ip").Str
			} else {
				resultC += value.Get("name").Str
			}
			// 端口
			resultC += ":" + sValue.Get("port").String()
			result = append(result, resultC)
			return true
		})
		return true
	})
	next = gResult.Get("links").Get("next").Str
	if next != "" {
		goto run
	}
	return
}
