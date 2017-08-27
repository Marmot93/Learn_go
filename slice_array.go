package main

import (
	"fmt"
)
// 数组是按值传递的（即是传递的副本），而切片是引用类型，传递切片的成本非常小，而且是不定长的。而且数组是定长的，而切片可以调整长度。
func main() {
	// int 型数组
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("len and cap of array %v is: %d and %d\n", a, len(a), cap(a))
	fmt.Printf("item in array: %v is:\n", a)
	// index, value
	for i, value := range a {
		fmt.Printf(" index: %d, value: %d\n", i, value)
	}

	s1 := a[3:6]
	fmt.Printf("len and cap of slice: %v is: %d and %d\n", s1, len(s1), cap(s1))
	fmt.Printf("item in slice: %v is:", s1)
	for _, value := range s1 {
		fmt.Printf("% d", value)
	}

	fmt.Println()

	s1[0] = 456
	// 切片是引用，所以会改变原来的值
	fmt.Printf("item 'a' in array changed after changing slice: %v is:", s1)
	for _, value := range a {
		fmt.Printf("% d", value)
	}

	fmt.Println()
	// 直接make一个slice
	s2 := make([]int, 10, 20)
	// index 4 change to  5
	s2[4] = 5
	fmt.Printf("len and cap of slice: %v is: %d and %d\n", s2, len(s2), cap(s2))
	fmt.Printf("item in slice %v is:", s2)
	for _, value := range s2 {
		fmt.Printf("% d", value)
	}

	fmt.Println()
}
