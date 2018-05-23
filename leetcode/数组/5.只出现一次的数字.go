package main

import (
	"fmt"
)

import "sort"

func singleNumber(nums []int) int {
	sort.Ints(nums)
	if len(nums) == 1 {
		return nums[0]
	}
	fmt.Println(nums)
	for len(nums) > 1 && nums[0] == nums[1] {
		nums = nums[2:]
	}
	return nums[0]
	//// 剔除相同的,直到只剩一个
	//for i := 1; len(nums) > 1 && i < len(nums); {
	//	if nums[i] == n {
	//		// 剔除后一个
	//		nums = append(nums[:i], nums[i+1:]...)
	//		// 剔除前一个
	//		nums = append(nums[:0], nums[1:]...)
	//		// n 置为 第一个
	//		n = nums[0]
	//		// 重置
	//		i = 1
	//	}
	//	i ++
	//}
	//return n
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}
