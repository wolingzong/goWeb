package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/didi/gendry/builder"
	"github.com/didi/gendry/manager"
	"github.com/didi/gendry/scanner"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dbName := "test"
	user := "root"
	password := "3dclobest"
	host := "3dcloserver"
	port := 3306

	var db *sql.DB
	var err error

	db, err = manager.
		New(dbName, user, password, host).
		Set(
			manager.SetCharset("utf8"),
			manager.SetAllowCleartextPasswords(true),
			manager.SetInterpolateParams(true),
			manager.SetTimeout(1*time.Second),
			manager.SetReadTimeout(1*time.Second)).
		Port(port).Open(true)

	where := map[string]interface{}{
		"city in":  []string{"beijing", "shanghai"},
		"score":    5,
		"age >":    35,
		"address":  builder.IsNotNull,
		"_orderby": "bonus desc",
		"_groupby": "department",
	}
	table := "person"
	selectFields := []string{"name", "age", "sex"}
	cond, values, err := builder.BuildSelect(table, where, selectFields)

	//cond = SELECT name,age,sex FROM some_table WHERE (score=? AND city IN (?,?) AND age>? AND address IS NOT NULL) GROUP BY department ORDER BY bonus DESC
	//values = []interface{}{"beijing", "shanghai", 5, 35}

	rows, err := db.Query(cond, values...)
	defer rows.Close()

	type Person struct {
		Name string `ddb:"name"`
		Age  int    `ddb:"age"`
		Sex  int    `ddb:"sex"`
	}

	var students []Person
	err = scanner.Scan(rows, &students)
	for _, student := range students {
		fmt.Println(student)

	}

	if nil != err {
		panic(err)
	}

}
