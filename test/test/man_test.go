package main_test

import "testing"

func TestWriteDirection(t *testing.T) {
	c := make(chan int, 100)
	onOnly(c)
	outOnly(c)
}

func onOnly(c chan<- int) {
	c <- 1
	// c <- 1 当c是单向写入channel时，不能再取数据。
}

func outOnly(c <-chan int) {
	<-c
	//<- c 当c是单向读取时，此时不能再写入数据。
}
