package main

import (
	"fmt"
	"math"
)

// 等待所有 goroutine 执行完毕
// 使用传址方式为 WaitGroup 变量传参
// 使用 channel 关闭 goroutine

func main() {
	s := fmt.Sprintf("%v", math.Pi)
	fmt.Println(s)
}
