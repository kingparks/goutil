# goutil
go语言工具

- 支持给字符数组洗牌
- 支持interface转化为具体类型
- 支持获取字符串中间字符

Installation
------------

Use go get.

	go get github.com/kingparks/goutil

Then import the validator package into your own code.

	import "github.com/kingparks/goutil"

示例:
```golang
// interface转化为string
fmt.Println(goutil.NewT(123).ToString())
// output:123

// 获取字符串中间字符
fmt.Println(goutil.Str("12345").GetBetween("1","4"))
// output:23

// 给字符数组洗牌
fmt.Println(goutil.StrArr{"1","2","3"}.Random2String())
// output:2,1,3

// 验证授权码
goutil.VerifyMachineLicense("serverName")
```


