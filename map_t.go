package main

import (
	"fmt"
)

func main() {
	// shiyanlou = {}
	shiyanlou := make(map[string]string) // 与 map[string]string 相同
	// shiyanlou["golang"] = "docker"
	shiyanlou["golang"] = "docker"
	shiyanlou["python"] = "flask web framework"
	shiyanlou["linux"] = "sys administrator"
	fmt.Print("Traverse all keys: ")
	// for key in shiyanlou.keys()
	for key := range shiyanlou {
		// print('{}'.format(key))
		fmt.Printf("% s ", key)
	}
	fmt.Println()
	// shiyanlou.delete('linux')
	delete(shiyanlou, "linux")
	shiyanlou["golang"] = "beego web framework"
	// value, key 真他妈的反人类，把值放前面，没有就是空
	v, key := shiyanlou["linux"]
	fmt.Printf("Found key \"linux\" Yes or False: %t, value of key \"linux\": \"%s\"", key, v)
	fmt.Println()

	fmt.Println("Traverse all keys/values after changed:")
	// for k,v in shiyanlou.items()
	for k, v := range shiyanlou { //遍历了映射的所有键/值对
		fmt.Printf("\"%s\": \"%s\"\n", k, v)
	}
}
