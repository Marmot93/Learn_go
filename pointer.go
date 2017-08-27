package main

import (
	"fmt"
)

func swap1(x, y, p *int) {
	if *x > *y {
		*x, *y = *y, *x
	}
	*p = *x * *y
}
// 第一个（）内是变量，第二个是输出
func swap2(x, y int) (int, int, int) {
	if x > y {
		x, y = y, x
	}
	return x, y, x * y
}

func main() {
	i := 9
	j := 5
	product := 0
	swap1(&i, &j, &product)
	// i，j 在内里面，找出地址输出即可，不需要函数 return
	fmt.Println(i, j, product)

	a := 64
	b := 23
	a, b, p := swap2(a, b)
	// 需要函数赋值然后 return
	fmt.Println(a, b, p)
}
