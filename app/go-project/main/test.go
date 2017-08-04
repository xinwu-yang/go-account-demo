package main

import (
	"fmt"
	"encoding/hex"
	"com.cxria/utils/crypto"
)

func main() {
	shaName := crypto.SHA256Hex("xinwuy")
	fmt.Println(shaName)
	k := []byte(shaName[0:32])
	i := []byte(shaName[32:64])
	hex.Decode(k, k)
	hex.Decode(i, i)
	x, _ := hex.DecodeString("90d44eeb5cd0ba3f6b84cb0d19a4f897")
	fmt.Println(x)
	y, _ := crypto.AesDecrypt(x, k[:16])
	fmt.Println(string(y))
}
