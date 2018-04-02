package main

import "fmt"

// 正确示例
func main() {
	s1 := []int{1,2,3}
	s2 := s1[:0]

	s2 = append(s2, 4)
	fmt.Print(s1)
	fmt.Print(s2)
}