package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"geek/pkg/apis"
	gobmi "github.com/armstrongli/go-bmi"
	_ "github.com/stretchr/testify/assert"
	_ "go/ast"
	_ "io"
	"log"
	"net"
	_ "time"
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

	rank := NewFatRateRank()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("建立连接失败：", err)
			continue
		}
		go talk(conn, rank)
	}
}

func talk(conn net.Conn, rank *FatRateRank) {
	defer fmt.Println("结束连接")
	defer conn.Close() //关闭连接
lab2:
	for {
		finalReceivedMessage := make([]byte, 0, 1024)
		encounterError := false
	lab1:
		for {
			buf := make([]byte, 1024)
			valid, err := conn.Read(buf) //读取byte的值，并得到读取到的byte的长度
			if err != nil {
				log.Println("读取数据出问题", err)
				log.Println("重新读取")
				encounterError = true
				break lab1
			}
			if valid == 0 {
				break lab1
			}
			content := buf[:valid]
			finalReceivedMessage = append(finalReceivedMessage, content...)
		}
		if encounterError {
			conn.Write([]byte("接收消息失败"))
			continue lab2
		}

		pi := &apis.PersonalInformation{}
		if err := json.Unmarshal(finalReceivedMessage, pi); err != nil {
			conn.Write([]byte("输入数据无法解析，重新录入"))
			continue lab2
		}
		bim, err := gobmi.BMI(float64(pi.Weight), float64(pi.Tall))
		if err != nil {
			conn.Write([]byte("计算体脂失败"))
			continue lab2
		}
		fr, err := calculator.CalcFatRate(bmi, pi.Sex)

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
