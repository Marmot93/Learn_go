package main

import (
	"fmt"
)

func main() {
	shiyanlou := make(map[string]string) // 与 map[string]string 相同
	shiyanlou["golang"] = "docker"
	shiyanlou["python"] = "flask web framework"
	shiyanlou["linux"] = "sys administrator"
	fmt.Print("Traverse all keys: ")
	for key := range shiyanlou { // 遍历了映射的所有键
		fmt.Printf("% s ", key)
	}
	fmt.Println()

	delete(shiyanlou, "linux")                  // 从映射中删除键"linux"及其值
	shiyanlou["golang"] = "beego web framework" // 更新键“golang"的值

	v, found := shiyanlou["linux"]
	fmt.Printf("Found key \"linux\" Yes or False: %t, value of key \"linux\": \"%s\"", found, v)
	fmt.Println()

	fmt.Println("Traverse all keys/values after changed:")
	for k, v := range shiyanlou { //遍历了映射的所有键/值对
		fmt.Printf("\"%s\": \"%s\"\n", k, v)
	}
}
