package main

import (
	"fmt"
	"strconv"
)

//一个回文数指的是从左向右和从右向左读都一样的数字。
//最大的由两个两位数乘积构成的回文数是 9009 = 91 * 99。
//找出最大的有由个三位数乘积构成的回文数。

func reverseString(s string) string {
	runes := []rune(s)

	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}

	return string(runes)
}

func main() {
	max := 0
	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			s := strconv.Itoa(i * j)
			if s == reverseString(s) && i*j > max {
				max = i * j
			}
		}
	}
	fmt.Print(max)
}
