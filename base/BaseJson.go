package base

import "encoding/json"

type Base struct {
	Ok         int
	ErrorCode  int
	Content    interface{}
	Array      interface{}
	ErrorArray interface{}
	Message    string
}

func (obj Base) String() string {
	returnData := make(map[string]interface{})
	returnData["b"] = obj.Ok
	if obj.ErrorCode != 0 {
		returnData["i"] = obj.ErrorCode
	}
	if obj.Array != nil {
		returnData["a"] = obj.Array
	}
	if obj.Content != nil {
		returnData["o"] = obj.Content
	}
	if obj.ErrorArray != nil {
		returnData["ec"] = obj.ErrorArray
	}
	b, _ := json.Marshal(returnData)
	return string(b)
}
