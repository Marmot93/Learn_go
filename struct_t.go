package main

import "fmt"

<<<<<<< HEAD
// 声明一个新的类型 class person()
type person struct {
	name string
	age int
}

// 比较两个人的年龄，返回年龄大的那个人，并且返回年龄差
// struct也是传值的
func Older(p1, p2 person) (person, int) {
	if p1.age>p2.age {  // 比较p1和p2这两个人的年龄
		return p1, p1.age-p2.age
	}
	return p2, p2.age-p1.age
}

func main() {
	// tom = person()
	var tom person

	// 赋值初始化
	tom.name, tom.age = "Tom", 18

	// 两个字段都写清楚的初始化m，这样比较好
	bob := person{age:25, name:"Bob"}

	// 按照struct定义顺序初始化值
	paul := person{"Paul", 43}

	tb_Older, tb_diff := Older(tom, bob)
	tp_Older, tp_diff := Older(tom, paul)
	bp_Older, bp_diff := Older(bob, paul)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, bob.name, tb_Older.name, tb_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		tom.name, paul.name, tp_Older.name, tp_diff)

	fmt.Printf("Of %s and %s, %s is older by %d years\n",
		bob.name, paul.name, bp_Older.name, bp_diff)
}
=======
// 使用"类型(type)"和其对应的"值"，其中自定义类型的值可以包含方法。
type Count int // 创建自定义类型 Count
// 给count添加方法
func (count *Count) Increment()  { *count++ }
func (count *Count) Decrement()  { *count-- }
func (count Count) IsZero() bool { return count == 0 }

type Part struct {
	// 基于结构体创建自定义类型 Part
	stat string
	Count // 匿名字段
}

func (part Part) IsZero() bool { // 覆盖了匿名字段Count的IsZero()方法
	return part.Count.IsZero() && part.stat == "" // 调用了匿名字段的方法
}

func (part Part) String() string { // 定义String()方法，自定义了格式化指令%v的输出
	return fmt.Sprintf("<<%s, %d>>", part.stat, part.Count)
}

func main() {
	var i Count = -1
	fmt.Printf("Start \"Count\" test:\nOrigin value of count: %d\n", i)
	i.Increment()
	fmt.Printf("Value of count after increment: %d\n", i)
	fmt.Printf("Count is zero t/f? : %t\n\n", i.IsZero())
	fmt.Println("Start: \"Part\" test:")
	part := Part{"232", 0}
	fmt.Printf("Part: %v\n", part)
	fmt.Printf("Part is zero t/f? : %t\n", part.IsZero())
	fmt.Printf("Count in Part is zero t/f?: %t\n", part.Count.IsZero()) // 尽管覆盖了匿名字段的方法，单还是可以访问
}
>>>>>>> origin/master
