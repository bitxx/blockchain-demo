package main

import "fmt"

//欧几里得算法，也就是辗转相除法

var (
	val1 int
	val2 int
	tmp  int
)

func main() {
	fmt.Scanln(&val1, &val2)
	for val2 != 0 {
		tmp = val2
		val2 = val1 % val2
		val1 = tmp
	}
	fmt.Printf("result:%d", val1)
}
