package main

import (
	"fmt"
	"net"
)

func Response(conn net.Conn) (err error) {
	buf := make([]byte, 1024)
	// 客户端有数据就读
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		fmt.Println("连接者：", conn.RemoteAddr().String(), "发送的内容：", string(buf))
		n, err = conn.Write([]byte(fmt.Sprintf("收到数据：%s", string(buf[:n]))))
		if err != nil {
			return err
		}
	}
}

func WriteToSever(conn net.Conn) (err error) {
	return
}
func main() {
	// 将addr作为TCP地址解析并返回，
	addr, err := net.ResolveTCPAddr("tcp", ":81")
	if err != nil {
		panic(err)
	}
	// 监听TCP地址，返回TCP监听器
	lsiten, err := net.ListenTCP("tcp", addr)
	if err != nil {
		panic(err)
	}
	fmt.Println("开始监听")
	// 用for循环持续监听客户端连接，有人连接就会产生一个conn对象
	// conn对象是TCP连接，可以读写数据
	for {
		conn, err := lsiten.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("有人连接", conn.RemoteAddr().String())
		// 向客户端发送数据
		conn.Write([]byte(fmt.Sprintf("你好,本服务器地址:%s, 你已成功连接!", lsiten.Addr().String())))
		go Response(conn)
	}

}
