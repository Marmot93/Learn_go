package main

import "fmt"

//问题： 13195 的质数因子有 5, 7, 13 和 29。
//600851475143 的最大质数因子是多少？

func main() {
	/*数字除以2，若能被整除，取结果继续除以 i = 2，若不能则除以 i+1：
	依序递增，直到被自己除，此时的数字即为所求的最大质因子。*/
	num := 600851475143
	i := 2
	for {
		if num == i {
			break
		}
		if num%i == 0 {
			num = num / i
		} else {
			i ++
		}
	}
	fmt.Print(num)
}
