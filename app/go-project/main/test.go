package main

//import (
//	"fmt"
//	"encoding/hex"
//	"com.cxria/utils/crypto"
//	"strings"
//	"com.cxria/modules/account/domain"
//)
//
//
//func main() {
//	shaName := crypto.SHA256Hex("19950304")
//	fmt.Println(shaName)
//	k := []byte(shaName[0:32])
//	i := []byte(shaName[32:64])
//	hex.Decode(k, k)
//	hex.Decode(i, i)
//	p,_ := crypto.AesEncrypt([]byte("杨欣武"), k[:16], i[:16])
//	fmt.Println(strings.ToUpper(hex.EncodeToString(p)))
//	myName := []byte("杨欣武")
//	fmt.Println(myName)
//	hexName := hex.EncodeToString(myName)
//	fmt.Println(hexName)
//
//	s := domain.Session{}
//	fmt.Println(s)
//}