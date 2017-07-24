package main

import (
	"fmt"
	"com.cxria/utils/crypto"
	"encoding/hex"
)

func main() {
	sha := crypto.SHA256Hex("18180423321")
	fmt.Println("SHA256_HEX :", sha)

	k := sha[0:32]
	i := sha[32:64]
	fmt.Println(k, i)

	key := []byte(k)
	iv := []byte(i)

	hex.Decode(key, key)
	hex.Decode(iv, iv)

	fmt.Println("key :", key, len(key))
	fmt.Println("iv :", iv, len(iv))

	str, _ := crypto.AesEncrypt([]byte("yxw123456"), key[:16], iv[:16])
	fmt.Println(hex.EncodeToString(str))
}
