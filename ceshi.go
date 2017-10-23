package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)




const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "blog"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
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

	var id,age int
	var name  string
	var create_time  time.Time
	rows, err := db.Query(`SELECT id, age, create_time, "name" FROM student;`)
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &age, &create_time, &name)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(id, nickname, create_time.Format("2006-01-02 15:04:05"))
		fmt.Println(id, age, name,create_time.Format("2006-01-02 15:04:05"))
	}

}
