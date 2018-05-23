package main

import "fmt"

func containsDuplicate(nums []int) bool {
	for i_index, i := range nums{
		for j_index, j := range nums{
			if i==j && i_index != j_index{
				return true
			}
		}
	}
	return false
}

func main() {
	nums := []int{1,2,3,1}
	x := containsDuplicate(nums)
	fmt.Print(x)
}