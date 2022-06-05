package main

import "fmt"

func main() {
	var a int = 1
	var b int = 2
	add(&a, &b)
	fmt.Println(a) //此时a的值为3
	c := &a
	d := &c
	fmt.Printf("a类型为：%T，a地址为：%v\n", &a, a) //a=*int
	fmt.Printf("c类型为：%T，c地址为：%v\n", &c, c) //c=**int
	fmt.Printf("d类型为：%T，d地址为：%v\n", &d, d) //d=***int

	e := add //e=*func(*int *int)
	fun1 := &e
	(*fun1)(&a, &b) //e=**func(*int *int)
	fmt.Printf("fun1类型为：%T，fun1地址为：%v", &fun1, fun1)

}

func add(a, b *int) {
	*a = *a + *b
}
