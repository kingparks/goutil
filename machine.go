package goutil

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"io"
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
	resp, err := httplib.Get("https://ghproxy.com/https://raw.githubusercontent.com/kingparks/license-key/main/" + appName + ".txt").DoRequest()
	if err != nil || resp.StatusCode != 200 {
		fmt.Println("ghproxy 失败", resp.StatusCode, err)
		resp, err = httplib.Get("https://raw.githubusercontent.com/kingparks/license-key/main/" + appName + ".txt").Retries(2).DoRequest()
		if err != nil || resp.StatusCode != 200 {
			fmt.Println("github 失败", resp.StatusCode, err)
			return false
		}
	}
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}
	macMD5 := getMacMD5()
	fmt.Println("设备码:", macMD5)
	if strings.Contains(string(bytes), macMD5) {
		return true
	}
	return false
}
