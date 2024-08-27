package wechat

import (
	"github.com/astaxie/beego/httplib"
	"github.com/kingparks/goutil/net"
	"github.com/tidwall/gjson"
	"time"
)

func getToken(corpID, agentID, corpSecret string) string {
	a, e := httplib.Get("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=" + corpID + "&corpsecret=" + corpSecret).String()
	if e != nil {
		return ""
	}
	return gjson.Get(a, "access_token").String()
}

func Send(corpID, agentID, corpSecret, msg string) {
	go func(corpID, agentID, corpSecret, msg string) {
		msg += time.Now().Format("\n\n2006-01-02 15:04")
		msg += "\n" + net.GetIPV4() + "\n"
		token := getToken(corpID, agentID, corpSecret)
		_, _ = httplib.Post("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=" + token).Body(`
{
   "touser" : "@all",
   "msgtype" : "text",
   "agentid" : ` + agentID + `,
   "text" : {
       "content" : "` + msg + `"
   },
   "enable_duplicate_check": 1,
}`).DoRequest()
	}(corpID, agentID, corpSecret, msg)
}
