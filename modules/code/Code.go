package code

var ErrorCode = make(map[string]int)
var Message = make(map[string]string)

func init() {
	var i int = 300
	i++
	ErrorCode["NO_SESSION"] = i
	Message["NO_SESSION"] = "登陆会话不存在"
	i++
}
