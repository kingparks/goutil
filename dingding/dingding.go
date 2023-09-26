package dingding

import (
	"encoding/json"
	"github.com/astaxie/beego/httplib"
)

func Send(token, text string) {
	go func(token, text string) {
		req := httplib.Post("https://oapi.dingtalk.com/robot/send?access_token=" + token)
		req.Header("Content-Type", "application/json")
		b, _ := json.Marshal(map[string]interface{}{
			"msgtype": "text",
			"text": map[string]interface{}{
				"content": text,
			},
		})
		_, _ = req.Body(b).DoRequest()
	}(token, text)
}
