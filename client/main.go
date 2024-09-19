package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("客户端启动")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("链接失败")
		return
	}
	fmt.Println("链接成功", conn)
}
