package util

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"net"
	"sort"
	"strings"
)

func getMacMD5() string {
	// 获取本机的MAC地址
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("err:", err)
		return ""
	}
	var macAddress []string
	for _, inter := range interfaces {
		if !strings.HasPrefix(inter.Name, "en") {
			continue
		}
		macAddress = append(macAddress, inter.HardwareAddr.String())
	}
	sort.Strings(macAddress)
	return fmt.Sprintf("%x", md5.Sum([]byte(strings.Join(macAddress, ","))))
}

// VerifyMachineLicense 验证服务授权
func VerifyMachineLicense(appName string) bool {
	res, err := httplib.Get("https://ghproxy.com/https://raw.githubusercontent.com/Jetereting/license-key/main/" + appName + ".txt").String()
	if err != nil {
		res, err = httplib.Get("https://raw.githubusercontent.com/Jetereting/license-key/main/" + appName + ".txt").Retries(7).String()
		if err != nil {
			return false
		}
	}
	macMD5 := getMacMD5()
	fmt.Println("设备码:", macMD5)
	if strings.Contains(res, macMD5) {
		return true
	}
	return false
}
