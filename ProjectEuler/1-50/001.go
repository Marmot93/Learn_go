package main

import "fmt"

//问题： 10 以下的自然数中，属于 3 或 5 的倍数的有 3, 5, 6 和 9，它们之和是 23,
//找出 1000 以下的自然数中，属于 3 或 5 的倍数的数字之和。

func main() {
	sum := 0
	for i := 1; i < 1000; i ++ {
		if i%5 == 0 || i%3 == 0 {
			sum += i
		}
	}
	fmt.Print(sum)
}
