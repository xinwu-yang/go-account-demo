package code

var ErrorCode = make(map[string]int)
var Message = make(map[string]string)

func init() {
	var i int = 300
	i++
	ErrorCode["NO_SESSION"] = i
	Message["NO_SESSION"] = "登陆会话不存在"
	i++
	ErrorCode["SMS_TIME_NOT_EX"] = i
	Message["SMS_TIME_NOT_EX"] = "验证码发送限制"
	i++
	ErrorCode["EXPIRED_CODE"] = i
	Message["EXPIRED_CODE"] = "验证码不存在或者已过期"
}
