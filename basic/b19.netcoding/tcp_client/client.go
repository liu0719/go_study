package main

import (
	"fmt"
	"net"
)

func Response(conn net.Conn) (err error) {
	var buf string
	// 发数据
	for {
		fmt.Println("输入内容,输入exit退出")
		fmt.Scanln(&buf)
		if buf == "exit" {
			break
		}
		_, err := conn.Write([]byte(buf))
		if err != nil {
			return err
		}
		fmt.Println("发送成功")
	}
	return
}

func main() {
	// dial 拨号，连接到服务器
	conn, err := net.Dial("tcp", ":81")
	if err != nil {
		panic(err)
	}
	// for循环
	for {
		go Response(conn)
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			panic(err)
		}
		fmt.Println("服务器", conn.RemoteAddr().String(), "发来：", string(buf[:n]))
	}

}
