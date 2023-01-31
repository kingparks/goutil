package goutil

import (
	"math/rand"
	"time"
)

type StrArr []string

// Random2String 打乱数组，返回一个字符串
func (sa StrArr) Random2String() (s string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(sa)) {
		val := sa[i]
		s += string(val) + ","
	}
	if len(s) > 0 {
		return s[:len(s)-1]
	}
	return
}

// Random2Arr 打乱数组
func (sa StrArr) Random2Arr() (s []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for _, i := range r.Perm(len(sa)) {
		val := sa[i]
		s = append(s, val)
	}
	return
}

// RemoveDuplicate 删除重复项
func (sa StrArr) RemoveDuplicate() []string {
	resArr := make([]string, 0)
	tmpMap := make(map[string]interface{})
	for _, val := range sa {
		//判断主键为val的map是否存在
		if _, ok := tmpMap[val]; !ok {
			resArr = append(resArr, val)
			tmpMap[val] = nil
		}
	}
	return resArr
}
