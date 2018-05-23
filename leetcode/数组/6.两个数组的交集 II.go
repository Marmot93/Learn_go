package main

import (
	"fmt"
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	s := nums1
	l := nums2
	if len(nums1) > len(nums2) {
		l, s = s, l
	}
	sort.Ints(s)
	sort.Ints(l)
	r := []int{}
	// 维护两个指针
	j, i := 0, 0
	for i < len(s) {
		for j < len(l) {
			if s[i] == l[j] {
				//值相同一起后移
				r = append(r, s[i])
				j ++
				break
			}
			// 内层移动
			j ++
		}
		i ++
	}
	return r
}

func main() {
	// 未通过
	nums1 := []int{1, 2}
	nums2 := []int{2, 1}
	fmt.Print(intersect(nums1, nums2))
}
