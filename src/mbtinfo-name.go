
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
	rows, err := db.Query("SELECT value FROM metadata where name='name'")

	checkErr(err)

	for rows.Next() {   
     	var value string   
    	err = rows.Scan(&value)
    	checkErr(err)
    	fmt.Println(value)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}