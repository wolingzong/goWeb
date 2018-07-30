package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//连接数据库

	db, err := sql.Open("mysql", "root:zhaoming8205@tcp(192.168.56.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	//查询
	age := 27
	rows, err := db.Query("SELECT name FROM person WHERE age=?", age)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%s is %d\n", name, age)
	}
	if err := rows.Err(); err != nil {
		log.Fatalln(err)
	}

}
