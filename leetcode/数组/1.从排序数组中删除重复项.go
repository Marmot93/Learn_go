package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	x := 0
	for i := 0; i < len(nums); {
		if nums[x] == nums[x+1] {
			// 剔除元素,保持不变
			nums = append(nums[:x], nums[x+1:]...)
		} else {
			x ++
			i ++
		}
		if x+1 == len(nums) {
			break
		}
	}
	return len(nums)
}

func main() {
	nums := []int{1,2,8,8,8,8,4}
	Len := removeDuplicates(nums)
	fmt.Print(Len)
}
