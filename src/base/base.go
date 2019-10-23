package base

import(
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
)

func GetInfo(mbtpath string) string {
	db, err := sql.Open("sqlite3", mbtpath)
	checkErr(err)

	//查询数据
	rows, err := db.Query("SELECT value FROM metadata where name='description'")
	checkErr(err)

	var value string 
	for rows.Next() {   
		err = rows.Scan(&value)
    	checkErr(err)    	
	}
	db.Close()
	return value
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}