package main

import "fmt"

func f(l []int, n int) [] int {
	if n > len(l)  {
		l = append(l[1:],l...)
		l[0] = 0
	}
	fmt.Println(l)
	if l[len(l)-n-1]+1 == 10 {
		l[len(l)-n-1] = 0
		return f(l, n+1)
	}

	l[len(l)-n-1] ++
	return l
}

func plusOne(digits []int) []int {
	if len(digits) == 1 {
		digits = f(digits, 0)
	}
	if digits[len(digits)-1]+1 == 10 {
		digits = f(digits, 0)
	} else {
		digits[len(digits)-1] ++
	}

	return digits
}

func main() {
	nums := []int{9}
	fmt.Println(plusOne(nums))

}
