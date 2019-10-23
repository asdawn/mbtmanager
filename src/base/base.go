package base

import(
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"strconv"
)

func Err(err int, msg string){
	switch err {
	case 1:
		fmt.Println("Invalid parameter: " + msg)
	case 2:
		fmt.Println ("File not exist: " + msg)
	case 3:
		fmt.Println ("Directory not exist: " + msg)
	case 4:
		fmt.Println ("Can not read file: " + msg)
	case 5:
		fmt.Println ("Can not read directory: " + msg)
	case 6:
		fmt.Println ("Can not write file: " + msg)
	case 7:
		fmt.Println ("Can not write directory: " + msg)
	default:
		fmt.Println ("Some error: " + msg)	
	}
}


func GetInfo(mbtpath string) string {
	db, err := sql.Open("sqlite3", mbtpath)
	checkErr(err)

	//查询数据
	rows, err := db.Query("SELECT name,value FROM metadata")
	checkErr(err)
	var name string
	var value string 
	var result string
	for rows.Next() {   
		err = rows.Scan(&name, &value)
		checkErr(err)
		result +=  name + ": " + value +"\n"	
	}
	db.Close()
	return result
}

func GetInfoWhich(mbtpath string) string {
	db, err := sql.Open("sqlite3", mbtpath)
	checkErr(err)
	var list [11]int
	//查询数据
	rows, err := db.Query("SELECT name FROM metadata")
	checkErr(err)
	var name string 
	for rows.Next() {   
		err = rows.Scan(&name)
		checkErr(err)
		switch name{
		case "name":
			list [0]++
		case "format":
			list[1]++
		case "bounds":
			list[2]++
		case "center":
			list[3]++
		case "minzoom":
			list[4]++
		case "maxzoom":
			list[4]++
		case "json":
			list[5]++
		case "attribution":
			list[6]++
		case "type":
			list[7]++
		case "version":
			list[8]++
		case "description":
			list[9]++
		default:
			list[10]++
		}	
	}
	db.Close()
	result := ""
	if list[0] == 1{
		result +="**name\n"
	}
	if list[1] == 1{
		result +="**format\n"
	}
	if list[5] == 1{
		result +="**json\n"
	}
	if list[2] == 1{
		result +="*bounds\n"
	}
	if list[3] == 1{
		result +="*center\n"
	}
	if list[4] == 2{
		result +="*minzoom and maxzoom\n"
	}
	if list[6] == 1{
		result +="attribution\n"
	}
	if list[7] == 1{
		result +="type\n"
	}
	if list[8] == 1{
		result +="version\n"
	}
	if list[9] == 1{
		result += "description\n"
	}
	if list[10] >= 1{
		result += "other: " + strconv.Itoa(list[10]) + "\n"
	}
	return result
}

func GetInfoField(mbtpath string, field string) string {
	db, err := sql.Open("sqlite3", mbtpath)
	checkErr(err)

	//查询属性
	rows, err := db.Query("SELECT value FROM metadata where name='" + field + "';")
	checkErr(err)

	var value string 
	rows.Next() 
	err = rows.Scan(&value)
    checkErr(err)    	
	db.Close()
	return value
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}