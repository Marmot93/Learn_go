package main

import "fmt"

// 冒泡排序
func maopao(list []int) {
	for i := 0; i < len(list); i++ {
		for j := i; j < len(list); j++ {
			if list[i] > list[j] {
				list[j], list[i] = list[i], list[j]
			}
		}
	}
}

// 鸡尾酒排序
func cocktail(list []int) {
	left := 0
	right := len(list) - 1
	for left < right {
		// 从左到右
		for i := left; i < right; i++ {
			if list[i] > list[i+1] {
				list[i+1], list[i] = list[i], list[i+1]
			}
		}
		right --
		// 从右到左
		for i := right; i > left; i-- {
			if list[i] < list[i-1] {
				list[i-1], list[i] = list[i], list[i-1]
			}
		}
		left ++
	}
}

//插入排序
func insertion(list []int) {
	for j := 1; j < len(list); j++ {
		// 把值拿出来
		temp := list[j]
		i := j - 1
		for ; i >= 0 && list[i] > temp; i-- {
			// 移动位置
			list[i+1] = list[i]
		}
		// 新值
		list[i+1] = temp
	}
}

//选择排序
func selection(list []int) {
	for i := 0; i < len(list); i++ {
		// 标记下标
		min := i
		for j := i + 1; j < len(list); j++ {
			if list[min] > list[j] {
				min = j
			}
		}
		// 交换位置
		list[min], list[i] = list[i], list[min]
	}
}

func main() {
	list := []int{11324, 12, 111, 118, 101, 70, 105, 115, 104, 67, 46, 99, 111, 109}
	fmt.Println(list)
	//maopao(list)
	//cocktail(list)
	//insertion(list)
	selection(list)
	fmt.Println(list)

}
