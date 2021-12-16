package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var catChan chan int
	catChan = make(chan int, 500)

	for i := 1; i <= 50; i++ {
		rand.Seed(time.Now().UnixNano())
		rand1 := rand.Intn(70)
		go write(rand1, catChan)
		go read(catChan)
	}
	defer func() {
		close(catChan)
	}()

	//for i := 0; i < 50; i++ {
	//	go read(catChan)
	//}

}

func write(w int, www chan int) {
	www <- w
	fmt.Println("写数据")
}

func read(cha chan int) {
	r := <-cha
	fmt.Println(r)
}
