package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
	checkErr(err)
	defer db.Close()
	sql1 := "SELECT id FROM test_table_2 WHERE relation_id = ?"
	stmt, err := db.Prepare(sql1)
	checkErr(err)
	res, err := stmt.Query("1")
	checkErr(err)

	for res.Next() {
		var id int
		err = res.Scan(&id)
		checkErr(err)
		fmt.Println("id=", id)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
