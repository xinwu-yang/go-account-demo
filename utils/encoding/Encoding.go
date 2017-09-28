package encoding

import (
	"encoding/base64"
	"github.com/djimenez/iconv-go"
)

func Base64Encode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func Base64Decode(src string) (string, error) {
	r, err := base64.StdEncoding.DecodeString(src)
	//utf8Subject := make([]byte, len(r))
	utf8Subject,err := iconv.ConvertString(string(r), "GB2312", "UTF-8")
	return string(utf8Subject), err
}
