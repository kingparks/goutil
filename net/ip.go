package net

import (
	"github.com/astaxie/beego/httplib"
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
