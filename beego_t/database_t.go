package main
import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"fmt"
	"time"
)

type barrecord struct {
	id int
	user_id int
	bar_id int
	date time.Location
	state int
}

func init() {
	// PostgreSQL 配置
	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
	orm.RegisterDataBase("default", "postgres", "user=postgres password=postgres dbname=chat host=127.0.0.1 port=5432 sslmode=disable")

	// 自动建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	barr := new(barrecord)
	barr.id=1
	o := orm.NewOrm()  //注册新的orm
	o.Read(barr)
	fmt.Println(barr)
}