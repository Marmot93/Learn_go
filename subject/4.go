package main

import "fmt"

type People struct{}

// showA() use showB()
func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

// showB()
func (p *People) ShowB() {
	fmt.Println("showB")
}


// 继承people
type Teacher struct {
	People
}

// 重写 showB 但是没有重写ShowA
func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}
