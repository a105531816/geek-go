package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("与server端建立连接失败！", err)
	}
	{
		defer conn.Close()
		for {
			r:=bufio.NewReader(os.Stdin)
			input,_,_ :=r.ReadLine()
			if len(input) !=0{
				call(conn, string(input))
			}

		}

	}
}

func call(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		log.Fatal("发送消息失败")
	} else {
		buf := make([]byte, 1024)
		valid, err := conn.Read(buf)
		if err != nil {
			log.Println("读数据失败！", err)
		} else {
			validData := buf[:valid]
			fmt.Printf("发送：%v，接收：%v\n", message, string(validData))
		}
	}
}
