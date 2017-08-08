package main

import (
	"net"
	"fmt"
	"bufio"
)

func main() {
	conn, err := net.Dial("tcp4", "pop.qq.com:995")
	if err != nil {
		fmt.Println("Dial : ", err)
		return
	}
	line,_,err := bufio.NewReader(conn).ReadLine()
	fmt.Println(string(line))
}
