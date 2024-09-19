package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(" 服务器启动")
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("链接失败:", err)
		return
	}
	for {
		conn, err2 := listen.Accept()
		if err != nil {
			fmt.Println("LOSE err2:", err2)
		} else {
			fmt.Println("等待链接成功,con=%v,客户端信息是%v/n", conn, conn.RemoteAddr().String())
		}

	}
}
