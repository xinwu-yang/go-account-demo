package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	pwd := []byte("yxw123456")
	hashedPassword, _ := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	fmt.Println(string(hashedPassword))
	fmt.Println(bcrypt.CompareHashAndPassword(hashedPassword,pwd))
}
