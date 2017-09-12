package foundation

import (
	"fmt"
)

func classchecker(items ...interface{}) { // 创建一个函数，该函数可以接受任意多的任意类型的参数
	// i index | x value
	for i, x := range items {
		switch x := x.(type) { // 创建了影子变量
		case bool:
			fmt.Printf("param #%d is a bool, value: %t\n", i, x)
		case float64:
			fmt.Printf("param #%d is a float64, value: %f\n", i, x)
		case int, int8, int16, int32, int64:
			fmt.Printf("param #%d is a int, value: %d\n", i, x)
		case uint, uint8, uint16, uint32, uint64:
			fmt.Printf("param #%d is a uint, value: %d\n", i, x)
		case nil:
			fmt.Printf("param #%d is a nil\n", i)
		case string:
			fmt.Printf("param #%d is a string, value: %s\n", i, x)
		default:
			fmt.Printf("param #%d is type is unknow\n", i)
		}
	}
}

func main() {
	classchecker(5, -17.98, "AIDEN", nil, true, complex(1, 1))

}