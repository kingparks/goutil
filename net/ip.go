package net

import (
	"github.com/astaxie/beego/httplib"
	"github.com/tidwall/gjson"
	"strings"
)

// GetIPV4 获取ipv4
func GetIPV4() string {
	ip, _ := httplib.Get("http://ipv4.ip.sb").Header("User-Agent", "curl").String()
	return strings.TrimSpace(ip)
}

// GetIP 获取ip
func GetIP() string {
	ip, _ := httplib.Get("http://ip.sb").Header("User-Agent", "curl").String()
	return strings.TrimSpace(ip)
}

// GetGEO 获取GEO
func GetGEO(ip ...string) string {
	geo, _ := httplib.Get("https://api.ip.sb/geoip/"+strings.Join(ip, ",")).
		Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36").
		String()
	gJsonGeo := gjson.Parse(geo)
	return gJsonGeo.Get("country").String() + "_" + gJsonGeo.Get("region").String() + "_" + gJsonGeo.Get("city").String()
}
