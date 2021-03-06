package main

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "blog"
)

type MainController struct {
	beego.Controller
}

func (self *MainController) Get() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	var id, age int
	var name string
	//var create_time time.Time
	rows, err := db.Query(`SELECT id, age, "name" FROM "user";`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &age, &name)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(id, nickname, create_time.Format("2006-01-02 15:04:05"))
		fmt.Println(id, age, name)
	}
	//self.Ctx.WriteString(fmt.Sprintf("id: %d age: %d name:%s", id, age, name))
	self.Ctx.WriteString("<h1>哪咤闹海</h1>")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
