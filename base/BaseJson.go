package base

import j "encoding/json"

type Json struct {
	Ok         int
	ErrorCode  int
	Content    interface{}
	Array      interface{}
	ErrorArray interface{}
	Message    string
}

func (json *Json) String() string {
	returnData := make(map[string]interface{})
	returnData["b"] = json.Ok
	if json.ErrorCode != 0 {
		returnData["i"] = json.ErrorCode
	}
	if json.Array != nil {
		returnData["a"] = json.Array
	}
	if json.Content != nil {
		returnData["o"] = json.Content
	}
	if json.ErrorArray != nil {
		returnData["ec"] = json.ErrorArray
	}
	b, _ := j.Marshal(returnData)
	return string(b)
}

func GetJson() Json {
	return Json{Ok: 0}
}
