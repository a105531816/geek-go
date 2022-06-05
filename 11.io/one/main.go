package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("/Users/mac/dat.txt")
	defer f.Close()
	if err != nil {
		fmt.Println("无法打开文件")
		os.Exit(2) //任意的非零错误码
	}
	b1 := make([]byte, 1024)
	n1, err := f.Read(b1)
	if err != nil {
		fmt.Println("无法读取文件")
		os.Exit(3) //任意的非零错误码
	}
	//必须使用string() 进行强转，否则输出byte的数字
	//必须使用字符串的截取进行处理，否则字符后面的0 0 0 0 就会被其他程序处理
	fmt.Printf("%s\n", string(b1[:n1]))
	fmt.Println("n1的大小", n1)
	f.Seek(0, io.SeekStart)   //seekstart表示从文件头开始进行offset
	f.Seek(0, io.SeekCurrent) //seekcurrent表示从目前游标位置，转移相对位置
	_, err = f.Write([]byte("fadsfasdfads\n"))
	if err != nil {
		fmt.Println("错误：", err)
		os.Exit(3) //任意的非零错误码
	}
	_, err = f.WriteString("writestri\n")
	if err != nil {
		fmt.Println("错误：", err)
		os.Exit(3) //任意的非零错误码
	}
	_, err = f.WriteAt([]byte("CHANGED"), 0)
	if err != nil {
		fmt.Println("错误：", err)
		os.Exit(3) //任意的非零错误码
	}
	f.Sync()
}
