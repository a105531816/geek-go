package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		//time.Now().UnixNano()
		rand.Seed(time.Now().UnixNano())
		fmt.Println(rand.Intn(10))
	}

}
func TestHest() {
	go fmt.Println("1")
}
