package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//顺序将值与随机位置的值进行交换
	poker := []int{51, 52, 53, 54, 55, 56, 57, 58, 59, 515, 511, 512, 513, 61, 62, 63, 64, 65, 66, 67, 68, 69, 616, 611, 612, 613, 71, 72, 73, 74, 75, 76, 77, 78, 79, 717, 711, 712, 713, 81, 82, 83, 84, 85, 86, 87, 88, 89, 818, 811, 812, 813, 100, 200}
	var b int
	for i := 0; i < len(poker); i++ {
		rand.Seed(time.Now().UnixNano())
		b = poker[i]
		c := rand.Intn(55)
		poker[i] = poker[c]
		poker[c] = b

	}
	fmt.Println(a)
}
