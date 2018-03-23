package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "param_record"
)

func main() {
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

	var id int
	var dailyTemperature string
	sqlString := `SELECT id, daily_temperature  FROM bms_service_forecast WHERE id=$1;`
	//rows, err := db.Query(`SELECT id, daily_temperature  FROM bms_service_forecast;`)

	stmt, err := db.Prepare(sqlString)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(3)

	for rows.Next() {
		err := rows.Scan(&id, &dailyTemperature)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(id, nickname, create_time.Format("2006-01-02 15:04:05"))
		fmt.Println(id, dailyTemperature)
	}

}
