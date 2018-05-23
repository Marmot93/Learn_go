package main

import "fmt"

func rotate(nums []int, k int) {
	for i:=0; i<k; i++ {
		x := nums[len(nums)-1]
		for index, i := range nums{
			nums[index] = x
			x = i
		}
	}
	fmt.Println(nums)
}

func main() {
	nums := []int{1,2,3,4,5,6,7}
	rotate(nums, 3)
	fmt.Print(nums)
}