package main

import (
	"fmt"
)

//冒泡排序
func main() {
	a := [...]int{7, 9, 3, 45, 6}
	fmt.Println(a)
	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++{
			if a[i] > a[j]{
				a[i],a[j] = a[j],a[i]
			}

		}
	}
	fmt.Println(a)
}
