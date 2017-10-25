package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq" //数据库驱动
)

type User struct {
	Id   int
	Age  int16
	Name string
}

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)

	orm.RegisterDataBase("default", "postgres", "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=blog sslmode=disable")

	orm.RegisterModel(new(User))
}

func insert() {

	orm.Debug = true;
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	user := User{Name: "marmot", Age: 24, Id:2}

	fmt.Println(o.Insert(&user))
}

func delete() {

	orm.Debug = true;
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	user := User{Id: 2}

	fmt.Println(o.Delete(&user))
}

func update() {

	orm.Debug = true;
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	user := User{Id: 8, Name: "koome_new"}

	fmt.Println(o.Update(&user))
}

func query() {

	var j, i int64

	orm.Debug = true;
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	var users [] User
	num, err := o.Raw("SELECT age, id, name FROM public.user").QueryRows(&users)
	fmt.Println(users)

	if err == nil {
		fmt.Println("user nums: ", num)
	}

	i = num
	for j = 0; j < i; j++ {
		fmt.Printf("Element[%d] : name:%s, age: %d\n", j, users[j].Name, users[j].Age)
	}
}

func main() {
	//insert()
	//delete()
	query()
}