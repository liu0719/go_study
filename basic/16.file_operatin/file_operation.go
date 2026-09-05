package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
)

// 文件操作

// 文件读

// 获取当前的go文件路径
func GetCurrentFilePath() string {
	_, filePath, _, _ := runtime.Caller(1)
	return filePath
}

func main() {
	// 一次性文件读取
	// 直接读取整个文件
	databyte, err := os.ReadFile("16.文件操作\\hello.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(databyte))

	// 获取当前go文件的路径
	fmt.Println("当前文件的路径为：", GetCurrentFilePath())

	// 文件太大就用open，只打开不读写
	file, err := os.Open("16.文件操作\\hello.txt")
	if err != nil {
		fmt.Println(err)
	}
	var dataByte = make([]byte, 13)
	for {
		n, err := file.Read(dataByte)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(dataByte[:n]))
	}
	file.Close()

	// 带缓冲读
	// 按行读取
	// 这里要重新打开，不然会从上一次停止的点继续读
	file, err = os.Open("16.文件操作\\hello.txt")
	if err != nil {
		fmt.Println(err)
	}
	buf := bufio.NewReader(file)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			fmt.Println("错误是：", err)
			break

		}

		fmt.Println("我是按行读的", string(line))

	}
	file.Close()

	// 按分隔符读取
	file, err = os.Open("16.文件操作\\hello.txt")
	if err != nil {
		fmt.Println(err)
	}

	// 实现read方法的结构体
	scanner := bufio.NewScanner(file)
	// 用该结构体的方法来指定分隔符读取文件
	scanner.Split(bufio.ScanWords) // 按行读取
	// scanner.Split(bufio.ScanLines) // 按单词读取
	// scanner.Split(bufio.ScanBytes) // 按字节读取
	// scanner.Split(bufio.ScanRunes) // 按字符读取

	// scanner.Scan()该方法会返回一个bool值，表示是否还有下一行
	for scanner.Scan() {
		// scanner.Text()就是读取到的内容
		fmt.Println("我是按分隔符读取的", scanner.Text())
	}
	file.Close()

	// 文件写入,创造和只写的方式
	// 第三个参数是权限,只在linux系统有用，第一位是用户权限，第二位是组权限，第三位是其他用户权限，4是读R，2是写W，1是执行X，相加就是对应用户的权限
	// 0444 表示三者均为只读的权限；
	// 0666 表示三者均为“读写”的权限；
	// 0777 表示三者均为读写执行的权限；
	// 0764 表示所有者有读写执行（7=4+2+1）的权限，组有读写（6=4+2）的权限，其他用户则为只读（4=4）；

	file, err = os.OpenFile("16.文件操作\\你好.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	file.Write([]byte("你好"))

	bytedata, err := io.ReadAll(file)
	if err != nil {
		//panic: read 16.文件操作\你好.txt: Access is denied. 读操作被拒绝
		// panic(err)
	}
	fmt.Println(string(bytedata))
	file.Close()

	// 全部写入
	// 直接写入整个文件
	err = os.WriteFile("16.文件操作\\一次全写.txt", []byte("我是一次性全部写入的文件内容"), 0666)
	if err != nil {
		panic(err)
	}

	// 文件拷贝
	// 读取文件
	rfile, err := os.Open("D:\\33542\\Pictures\\塔菲.webp")
	if err != nil {
		panic(err)
	}
	// 写入文件
	wfile, err := os.OpenFile("16.文件操作\\塔菲.webp", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	// 拷贝文件
	io.Copy(wfile, rfile)
	rfile.Close()
	wfile.Close()

	// 读取目录
	dir, err := os.ReadDir("16.文件操作")
	if err != nil {
		panic(err)
	}
	for _, entry := range dir {
		info, err := entry.Info()
		if err != nil {
			panic(err)
		}
		fmt.Println(entry.IsDir(), entry.Name(), info.Size())
	}
}
