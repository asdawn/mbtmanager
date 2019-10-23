
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "/mnt/c/tmp/blsub.mbtiles")
	checkErr(err)

	//查询数据
	rows, err := db.Query("SELECT * FROM metadata")

	checkErr(err)

	for rows.Next() {   
    	var name string
    	var value string   
    	err = rows.Scan(&name, &value)
    	checkErr(err)
    	fmt.Print(name+" --> ")
    	fmt.Println(value)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}