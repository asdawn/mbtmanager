package base

import(
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	"fmt"
	"strconv"
	"io/ioutil"
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

func GetInfoStatistics(mbtpath string) string {
	db, err := sql.Open("sqlite3", mbtpath)
	checkErr(err)
	zlist := make([]int, 0, 20)

	rows, err := db.Query("SELECT distinct zoom_level FROM tiles order by zoom_level")
	checkErr(err)
	var z int 
	for rows.Next() {   
		err = rows.Scan(&z)
		checkErr(err)
		zlist = append(zlist, z)
	}
	result := "z\tminx\tmaxx\tminy\tmaxy\ttilecount\n"
	for _, zz := range zlist {
		var minx, maxx, miny, maxy int
		var tilecount int64
		rows, err = db.Query("SELECT min(tile_column), max(tile_column), min(tile_row), max(tile_row), count(1) FROM tiles where zoom_level=" + strconv.Itoa(zz))
		checkErr(err)
		rows.Next()
		err = rows.Scan(&minx, &maxx, &miny, &maxy, &tilecount)
		checkErr(err)
		s := fmt.Sprintf("%d\t%d\t%d\t%d\t%d\t%d\n", zz, minx, maxx, miny, maxy, tilecount)
		result += s;
    }
	db.Close()
	return result
}

func GetInfoField(mbtpath string, field string) string {
	db, err := sql.Open("sqlite3", mbtpath)
	checkErr(err)

	rows, err := db.Query("SELECT value FROM metadata where name='" + field + "';")
	checkErr(err)

	var value string 
	rows.Next() 
	err = rows.Scan(&value)
    checkErr(err)    	
	db.Close()
	return value
}

func GetTile(mbtpath string, z int, x int, y int) []byte {
	db, err := sql.Open("sqlite3", mbtpath)
	checkErr(err)

	rows, err := db.Query("select tile_data from tiles where zoom_level = ? and tile_column = ? and tile_row = ? limit 1", z, x, y)
	checkErr(err)

	var tiledata []byte
	if rows.Next(){ 
		err = rows.Scan(&tiledata)
		checkErr(err)
		db.Close()    	
		return tiledata	
	}else{
		db.Close()
		return nil
	}
}

func GetTileBatch(sql.DB db, z int, x int, y int, tiledata [] byte) []byte {

	rows, err := db.Query("select tile_data from tiles where zoom_level = ? and tile_column = ? and tile_row = ? limit 1", z, x, y)
	checkErr(err)

	if rows.Next(){ 
		err = rows.Scan(&tiledata)
		checkErr(err)
		return tiledata	
	}else{
		return nil
	}
}

func SetTileBatch(sql.DB db, tiledata []byte, z int, x int, y int) error{


}

//todo
func SetTile(mbtpath string, tiledata []byte, z int, x int, y int)  error {
	db, err := sql.Open("sqlite3", mbtpath)
	checkErr(err)

	rows, err := db.Query("select tile_data from tiles where zoom_level = ? and tile_column = ? and tile_row = ? limit 1", z, x, y)
	checkErr(err)

	var tiledata []byte
	if rows.Next(){ 
		err = rows.Scan(&tiledata)
		checkErr(err)
		db.Close()    	
		return tiledata	
	}else{
		db.Close()
		return nil
	}
}

stmt, err := db.Prepare("insert into user(name,age)values(?,?)")
if err != nil {
    log.Println(err)
}

func ReadTileFile() []byte{
	return nil;
}

func WriteTileFile(tilepath string, tiledata []byte){
	//0664: rw, rw, r
	err := ioutil.WriteFile(tilepath, tiledata, 0664)
	checkErr(err)
}

/////todo
func AssureZoom(mbtpath string, z int) []byte {
	db, err := sql.Open("sqlite3", mbtpath)
	checkErr(err)

	rows, err := db.Query("select tile_data from tiles where zoom_level = ? and tile_column = ? and tile_row = ? limit 1", z, x, y)
	checkErr(err)

	var tiledata []byte
	if rows.Next(){ 
		err = rows.Scan(&tiledata)
		checkErr(err)
		db.Close()    	
		return tiledata	
	}else{
		db.Close()
		return nil
	}
}

func AssureIndex(mbtpath string) {
	db, err := sql.Open("sqlite3", mbtpath)
	checkErr(err)

	fmt.Println("Create index if not exists tileindex on tiles (zoom_level, tile_column, tile_row)")
	err := db.Execute("Create index if not exists tileindex on tiles (zoom_level, tile_column, tile_row")
}




func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}