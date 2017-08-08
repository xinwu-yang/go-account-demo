package sms

import (
	"time"
	"com.cxria/utils/crypto"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/astaxie/beego/logs"
)

const (
	SID     string = "05bfa8a845739e5e1117783ff0f1aabf"
	APP_ID  string = "803f148bb3974b9b9e5ef01ec9cd63da"
	TOKEN   string = "d42a5bb98b7fca35c227049d45c01d03"
	OVERDUE string = "15"
)

func sign(t string) string {
	return crypto.MD5(SID + t + TOKEN)
}

func SendByYzx(mobile, code string) bool {
	now := time.Now().Format("20060102150405000")
	var buf bytes.Buffer
	buf.WriteString("https://www.ucpaas.com/maap/sms/code")
	buf.WriteString("?sid=")
	buf.WriteString(SID)
	buf.WriteString("&appId=")
	buf.WriteString(APP_ID)
	buf.WriteString("&sign=")
	buf.WriteString(sign(now))
	buf.WriteString("&time=")
	buf.WriteString(now)
	buf.WriteString("&to=")
	buf.WriteString(mobile)
	buf.WriteString("&param=")
	buf.WriteString(code)
	buf.WriteString(",")
	buf.WriteString(OVERDUE)
	buf.WriteString("&templateId=")
	if len(mobile) == 11 {
		buf.WriteString("13831")
	} else {
		buf.WriteString("32650")
	}
	resp, _ := http.Get(buf.String())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var resultMap map[string]map[string]string
	json.Unmarshal(body, &resultMap)
	if resultMap["resp"]["respCode"] == "000000" {
		return true
	}
	logs.Error(string(body))
	return false
}
