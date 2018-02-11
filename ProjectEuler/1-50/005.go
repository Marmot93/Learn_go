package main

import (
	"fmt"
)

//2520 是最小的能被 1-10 中每个数字整除的正整数。
//最小的能被 1-20 中每个数整除的正整数是多少？

func main() {
	num := 2 ^ 4*3 ^ 2*5*7*11*13*17*19
	fmt.Println(num)
}
