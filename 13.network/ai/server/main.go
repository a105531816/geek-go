package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var qa = map[string]string{
	"你好":       "你好",
	"你是谁":      "我是小s",
	"你是男是女":    "你猜猜看",
	"今天天启怎么样？": "今天天气不错",
	"再见":       "再见",
}

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "nil")
	flag.Parse()
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("启动监听程序失败！", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("建立连接失败：", err)
			continue
		}
		go talk(conn)
	}
}

func talk(conn net.Conn) {
	defer fmt.Println("结束连接")
	defer conn.Close() //关闭连接
	for {
		buf := make([]byte, 1024)
		valid, err := conn.Read(buf) //读取byte的值，并得到读取到的byte的长度
		if err != nil {
			if err == io.EOF {
				time.Sleep(1 * time.Second)
				continue
			}
			log.Println("读数据失败！", err)
			continue
		}
		content := buf[:valid]          //获取buf中的有效值[:endkey]
		resp, ok := qa[string(content)] //将byte转成string即可通过key得到value
		if !ok {
			log.Println("没有听到回答")
			conn.Write([]byte("无法听懂"))
			continue
		}
		conn.Write([]byte(resp))
		if string(content) == "再见" {
			break
		}
	}
}
