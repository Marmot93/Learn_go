package main

import (
	"fmt"
	"reflect"
)

const a = 2 >> 1

func main() {
	fmt.Print(reflect.TypeOf(a))
}
